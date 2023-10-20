package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInterfaceLif() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif`: Logical interface\n\n",
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
							ValidateFunc: validation.IntBetween(1, 199),
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list (Named Access List)",
							ValidateFunc: validation.StringLenBetween(1, 16),
						},
					},
				},
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
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
										ValidateFunc: validation.IntBetween(0, 255),
									},
									"method": {
										Type: schema.TypeString, Optional: true, Description: "'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;",
										ValidateFunc: validation.StringInSlice([]string{"md5", "meticulous-md5", "meticulous-sha1", "sha1", "simple"}, false),
									},
									"password": {
										Type: schema.TypeString, Optional: true, Description: "Key String",
									},
									//omit encrypted field: 'encrypted'
								},
							},
						},
						"echo": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD Echo",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"demand": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Demand mode",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"interval_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type: schema.TypeInt, Optional: true, Description: "Transmit interval between BFD packets (Milliseconds)",
										ValidateFunc: validation.IntBetween(48, 1000),
									},
									"min_rx": {
										Type: schema.TypeInt, Optional: true, Description: "Minimum receive interval capability (Milliseconds)",
										ValidateFunc: validation.IntBetween(48, 1000),
									},
									"multiplier": {
										Type: schema.TypeInt, Optional: true, Description: "Multiplier value used to compute holddown (value used to multiply the interval)",
										ValidateFunc: validation.IntBetween(3, 50),
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
			"ifname": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Lif interface name",
				ValidateFunc: validation.StringLenBetween(1, 15),
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dhcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"address_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
										ValidateFunc: validation.IsIPv4Address,
									},
									"ipv4_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
										ValidateFunc: axapi.IsIpv4NetmaskBrief,
									},
								},
							},
						},
						"allow_promiscuous_vip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow traffic to be associated with promiscuous VIP",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"cache_spoofing_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "This interface connects to spoofing cache",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"generate_membership_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Membership Query",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"query_interval": {
							Type: schema.TypeInt, Optional: true, Default: 125, Description: "1 - 255 (Default is 125)",
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"max_resp_time": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Maximum Response Time (Max Response Time (Default is 100))",
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"unnumbered": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface as unnumbered",
							ValidateFunc: validation.IntBetween(0, 1),
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
													ValidateFunc: validation.StringLenBetween(1, 128),
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
																ValidateFunc: validation.StringLenBetween(1, 16),
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
																ValidateFunc: validation.StringInSlice([]string{"md5", "text"}, false),
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
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
														},
													},
												},
											},
										},
									},
									"send_packet": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable sending packets through the specified interface",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"receive_packet": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable receiving packet through the specified interface",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"send_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement transmission",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"version": {
													Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-compatible': RIPv1-compatible; '1-2': RIP version 1 & 2;",
													ValidateFunc: validation.StringInSlice([]string{"1", "2", "1-compatible", "1-2"}, false),
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
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"version": {
													Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;",
													ValidateFunc: validation.StringInSlice([]string{"1", "2", "1-2"}, false),
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
													ValidateFunc: validation.StringInSlice([]string{"poisoned", "disable", "enable"}, false),
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
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"value": {
																Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
																ValidateFunc: validation.StringInSlice([]string{"message-digest", "null"}, false),
															},
														},
													},
												},
												"authentication_key": {
													Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
													ValidateFunc: validation.StringLenBetween(1, 8),
												},
												"bfd_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bfd": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"disable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
																ValidateFunc: validation.IntBetween(0, 1),
															},
														},
													},
												},
												"cost": {
													Type: schema.TypeInt, Optional: true, Description: "Interface cost",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"database_filter_cfg": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"database_filter": {
																Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
																ValidateFunc: validation.StringInSlice([]string{"all"}, false),
															},
															"out": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
																ValidateFunc: validation.IntBetween(0, 1),
															},
														},
													},
												},
												"dead_interval": {
													Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"disable": {
													Type: schema.TypeString, Optional: true, Description: "'all': All functionality;",
													ValidateFunc: validation.StringInSlice([]string{"all"}, false),
												},
												"hello_interval": {
													Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"message_digest_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message_digest_key": {
																Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"md5": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"md5_value": {
																			Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
																		},
																		//omit encrypted field: 'encrypted'
																	},
																},
															},
														},
													},
												},
												"mtu": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF interface MTU (MTU size)",
													ValidateFunc: validation.IntBetween(576, 65535),
												},
												"mtu_ignore": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"network": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"broadcast": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF broadcast multi-access network",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"non_broadcast": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF NBMA network",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"point_to_point": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-point network",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"point_to_multipoint": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-multipoint network",
																ValidateFunc: validation.IntBetween(0, 1),
															},
															"p2mp_nbma": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
																ValidateFunc: validation.IntBetween(0, 1),
															},
														},
													},
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
													ValidateFunc: validation.IntBetween(0, 255),
												},
												"retransmit_interval": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"transmit_delay": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
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
													ValidateFunc: validation.IsIPv4Address,
												},
												"authentication": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
													ValidateFunc: validation.StringInSlice([]string{"message-digest", "null"}, false),
												},
												"authentication_key": {
													Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
													ValidateFunc: validation.StringLenBetween(1, 8),
												},
												"cost": {
													Type: schema.TypeInt, Optional: true, Description: "Interface cost",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"database_filter": {
													Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
													ValidateFunc: validation.StringInSlice([]string{"all"}, false),
												},
												"out": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"dead_interval": {
													Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"hello_interval": {
													Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"message_digest_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message_digest_key": {
																Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"md5_value": {
																Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
															},
															//omit encrypted field: 'encrypted'
														},
													},
												},
												"mtu_ignore": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
													ValidateFunc: validation.IntBetween(0, 255),
												},
												"retransmit_interval": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"transmit_delay": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
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
										ValidateFunc: axapi.IsIpv6AddressPlen,
									},
									"anycast": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 anycast address",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"link_local": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 link local address",
										ValidateFunc: validation.IntBetween(0, 1),
									},
								},
							},
						},
						"ipv6_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
							ValidateFunc: validation.IntBetween(0, 1),
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
													ValidateFunc: validation.IntBetween(0, 1),
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
																ValidateFunc: validation.IntBetween(0, 4294967295),
															},
															"area_id_addr": {
																Type: schema.TypeString, Optional: true, Description: "OSPF area ID in IP address format",
																ValidateFunc: validation.IsIPv4Address,
															},
															"tag": {
																Type: schema.TypeString, Optional: true, Description: "Set the OSPFv3 process tag",
																ValidateFunc: validation.StringLenBetween(1, 128),
															},
															"instance_id": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface instance ID",
																ValidateFunc: validation.IntBetween(0, 255),
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
													Type: schema.TypeString, ForceNew: true, Optional: true, Description: "ISO routing area tag",
													ValidateFunc: validation.StringLenBetween(1, 128),
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
													ValidateFunc: validation.StringInSlice([]string{"broadcast", "non-broadcast", "point-to-point", "point-to-multipoint"}, false),
												},
												"p2mp_nbma": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"network_instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
												},
											},
										},
									},
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"cost_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cost": {
													Type: schema.TypeInt, Optional: true, Description: "Interface cost",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IsIPv6Address,
												},
												"neig_inst": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
												},
												"neighbor_cost": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF cost for point-to-multipoint neighbor (metric)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"neighbor_poll_interval": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF dead-router polling interval (Seconds)",
													ValidateFunc: validation.IntBetween(0, 4294967295),
												},
												"neighbor_priority": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF priority of non-broadcast neighbor",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(0, 255),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
													ValidateFunc: validation.IntBetween(0, 255),
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
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication send-only for level-1 PDUs; 'level-2': Specify authentication send-only for level-2 PDUs;",
													ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
													ValidateFunc: validation.StringInSlice([]string{"md5"}, false),
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication mode for level-1 PDUs; 'level-2': Specify authentication mode for level-2 PDUs;",
													ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication key-chain for level-1 PDUs; 'level-2': Specify authentication key-chain for level-2 PDUs;",
													ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"disable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
										ValidateFunc: validation.IntBetween(0, 1),
									},
								},
							},
						},
						"circuit_type": {
							Type: schema.TypeString, Optional: true, Default: "level-1-2", Description: "'level-1': Level-1 only adjacencies are formed; 'level-1-2': Level-1-2 adjacencies are formed; 'level-2-only': Level-2 only adjacencies are formed;",
							ValidateFunc: validation.StringInSlice([]string{"level-1", "level-1-2", "level-2-only"}, false),
						},
						"csnp_interval_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"csnp_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Set CSNP interval in seconds (CSNP interval value)",
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Speficy interval for level-1 CSNPs; 'level-2': Specify interval for level-2 CSNPs;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
									},
								},
							},
						},
						"padding": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Add padding to IS-IS hello packets",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"hello_interval_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Set Hello interval in seconds (Hello interval value)",
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
										ValidateFunc: validation.IntBetween(2, 100),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello multiplier for level-1 IIHs; 'level-2': Specify hello multiplier for level-2 IIHs;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
									},
								},
							},
						},
						"lsp_interval": {
							Type: schema.TypeInt, Optional: true, Default: 33, Description: "Set LSP transmission interval (LSP transmission interval (milliseconds))",
							ValidateFunc: validation.IntBetween(1, 4294967295),
						},
						"mesh_group": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Mesh group number",
										ValidateFunc:  validation.IntBetween(1, 4294967295),
										ConflictsWith: []string{"isis.0.mesh_group.0.blocked"},
									},
									"blocked": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Block LSPs on this interface",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"isis.0.mesh_group.0.value"},
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
										ValidateFunc: validation.IntBetween(1, 63),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
									},
								},
							},
						},
						"network": {
							Type: schema.TypeString, Optional: true, Description: "'broadcast': Specify IS-IS broadcast multi-access network; 'point-to-point': Specify IS-IS point-to-point network;",
							ValidateFunc: validation.StringInSlice([]string{"broadcast", "point-to-point"}, false),
						},
						"password_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type: schema.TypeString, Optional: true, Description: "Configure the authentication password for interface",
										ValidateFunc: validation.StringLenBetween(1, 254),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify password for level-1 PDUs; 'level-2': Specify password for level-2 PDUs;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
										ValidateFunc: validation.IntBetween(0, 127),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify priority for level-1 routing; 'level-2': Specify priority for level-2 routing;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
									},
								},
							},
						},
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set per-LSP retransmission interval (Interval between retransmissions of the same LSP (seconds))",
							ValidateFunc: validation.IntBetween(0, 65535),
						},
						"wide_metric_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"wide_metric": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the wide metric for interface",
										ValidateFunc: validation.IntBetween(1, 16777214),
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;",
										ValidateFunc: validation.StringInSlice([]string{"level-1", "level-2"}, false),
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
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num_pkts': some help string; 'num_total_bytes': some help string; 'num_unicast_pkts': some help string; 'num_broadcast_pkts': some help string; 'num_multicast_pkts': some help string; 'num_tx_pkts': some help string; 'num_total_tx_bytes': some help string; 'num_unicast_tx_pkts': some help string; 'num_broadcast_tx_pkts': some help string; 'num_multicast_tx_pkts': some help string; 'dropped_dis_rx_pkts': some help string; 'dropped_rx_pkts': some help string; 'dropped_dis_tx_pkts': some help string; 'dropped_tx_pkts': some help string; 'dropped_rx_pkts_gre_key': some help string;",
							ValidateFunc: validation.StringInSlice([]string{"all", "num_pkts", "num_total_bytes", "num_unicast_pkts", "num_broadcast_pkts", "num_multicast_pkts", "num_tx_pkts", "num_total_tx_bytes", "num_unicast_tx_pkts", "num_broadcast_tx_pkts", "num_multicast_tx_pkts", "dropped_dis_rx_pkts", "dropped_rx_pkts", "dropped_dis_tx_pkts", "dropped_tx_pkts", "dropped_rx_pkts_gre_key"}, false),
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
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

func getObjectInterfaceLifAccessList(d []interface{}) edpt.InterfaceLifAccessList {
	var ret edpt.InterfaceLifAccessList
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceLifBfd(d []interface{}) edpt.InterfaceLifBfd {
	var ret edpt.InterfaceLifBfd
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifBfdAuthentication(in["authentication"].([]interface{}))
		ret.Echo = in["echo"].(int)
		ret.Demand = in["demand"].(int)
		ret.IntervalCfg = getObjectInterfaceLifBfdIntervalCfg(in["interval_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifBfdAuthentication(d []interface{}) edpt.InterfaceLifBfdAuthentication {
	var ret edpt.InterfaceLifBfdAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifBfdIntervalCfg(d []interface{}) edpt.InterfaceLifBfdIntervalCfg {
	var ret edpt.InterfaceLifBfdIntervalCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func getObjectInterfaceLifIp(d []interface{}) edpt.InterfaceLifIp {
	var ret edpt.InterfaceLifIp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.AddressList = getSliceInterfaceLifIpAddressList(in["address_list"].([]interface{}))
		ret.AllowPromiscuousVip = in["allow_promiscuous_vip"].(int)
		ret.CacheSpoofingPort = in["cache_spoofing_port"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.QueryInterval = in["query_interval"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		ret.Unnumbered = in["unnumbered"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceLifIpRouter(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceLifIpRip(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpOspf(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIpAddressList(d []interface{}) []edpt.InterfaceLifIpAddressList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpAddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpRouter(d []interface{}) edpt.InterfaceLifIpRouter {
	var ret edpt.InterfaceLifIpRouter
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Isis = getObjectInterfaceLifIpRouterIsis(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRouterIsis(d []interface{}) edpt.InterfaceLifIpRouterIsis {
	var ret edpt.InterfaceLifIpRouterIsis
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpRip(d []interface{}) edpt.InterfaceLifIpRip {
	var ret edpt.InterfaceLifIpRip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifIpRipAuthentication(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceLifIpRipSendCfg(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceLifIpRipReceiveCfg(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceLifIpRipSplitHorizonCfg(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthentication(d []interface{}) edpt.InterfaceLifIpRipAuthentication {
	var ret edpt.InterfaceLifIpRipAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Str = getObjectInterfaceLifIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLifIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLifIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationStr(d []interface{}) edpt.InterfaceLifIpRipAuthenticationStr {
	var ret edpt.InterfaceLifIpRipAuthenticationStr
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationMode(d []interface{}) edpt.InterfaceLifIpRipAuthenticationMode {
	var ret edpt.InterfaceLifIpRipAuthenticationMode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceLifIpRipAuthenticationKeyChain {
	var ret edpt.InterfaceLifIpRipAuthenticationKeyChain
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSendCfg(d []interface{}) edpt.InterfaceLifIpRipSendCfg {
	var ret edpt.InterfaceLifIpRipSendCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipReceiveCfg(d []interface{}) edpt.InterfaceLifIpRipReceiveCfg {
	var ret edpt.InterfaceLifIpRipReceiveCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceLifIpRipSplitHorizonCfg {
	var ret edpt.InterfaceLifIpRipSplitHorizonCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpOspf(d []interface{}) edpt.InterfaceLifIpOspf {
	var ret edpt.InterfaceLifIpOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceLifIpOspfOspfGlobal(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceLifIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobal(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobal {
	var ret edpt.InterfaceLifIpOspfOspfGlobal
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceLifIpOspfOspfGlobalBfdCfg(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceLifIpOspfOspfGlobalNetwork(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg {
	var ret edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalBfdCfg(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalBfdCfg {
	var ret edpt.InterfaceLifIpOspfOspfGlobalBfdCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg {
	var ret edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg(d []interface{}) []edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5 {
	var ret edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalNetwork(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalNetwork {
	var ret edpt.InterfaceLifIpOspfOspfGlobalNetwork
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Broadcast = in["broadcast"].(int)
		ret.NonBroadcast = in["non_broadcast"].(int)
		ret.PointToPoint = in["point_to_point"].(int)
		ret.PointToMultipoint = in["point_to_multipoint"].(int)
		ret.P2mpNbma = in["p2mp_nbma"].(int)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfIpList(d []interface{}) []edpt.InterfaceLifIpOspfOspfIpList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfIpList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6(d []interface{}) edpt.InterfaceLifIpv6 {
	var ret edpt.InterfaceLifIpv6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AddressList = getSliceInterfaceLifIpv6AddressList(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceLifIpv6Router(in["router"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpv6Ospf(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIpv6AddressList(d []interface{}) []edpt.InterfaceLifIpv6AddressList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6AddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Anycast = in["anycast"].(int)
		oi.LinkLocal = in["link_local"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6Router(d []interface{}) edpt.InterfaceLifIpv6Router {
	var ret edpt.InterfaceLifIpv6Router
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ripng = getObjectInterfaceLifIpv6RouterRipng(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpv6RouterOspf(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceLifIpv6RouterIsis(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterRipng(d []interface{}) edpt.InterfaceLifIpv6RouterRipng {
	var ret edpt.InterfaceLifIpv6RouterRipng
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterOspf(d []interface{}) edpt.InterfaceLifIpv6RouterOspf {
	var ret edpt.InterfaceLifIpv6RouterOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AreaList = getSliceInterfaceLifIpv6RouterOspfAreaList(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLifIpv6RouterOspfAreaList(d []interface{}) []edpt.InterfaceLifIpv6RouterOspfAreaList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6RouterOspfAreaList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6RouterOspfAreaList
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterIsis(d []interface{}) edpt.InterfaceLifIpv6RouterIsis {
	var ret edpt.InterfaceLifIpv6RouterIsis
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpv6Ospf(d []interface{}) edpt.InterfaceLifIpv6Ospf {
	var ret edpt.InterfaceLifIpv6Ospf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NetworkList = getSliceInterfaceLifIpv6OspfNetworkList(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceLifIpv6OspfCostCfg(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceLifIpv6OspfDeadIntervalCfg(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceLifIpv6OspfHelloIntervalCfg(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceLifIpv6OspfMtuIgnoreCfg(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceLifIpv6OspfNeighborCfg(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceLifIpv6OspfPriorityCfg(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceLifIpv6OspfTransmitDelayCfg(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNetworkList(d []interface{}) []edpt.InterfaceLifIpv6OspfNetworkList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNetworkList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNetworkList
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfCostCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfCostCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfDeadIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfDeadIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfHelloIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfHelloIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfMtuIgnoreCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfMtuIgnoreCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNeighborCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfNeighborCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNeighborCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNeighborCfg
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfPriorityCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfPriorityCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfTransmitDelayCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfTransmitDelayCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsis(d []interface{}) edpt.InterfaceLifIsis {
	var ret edpt.InterfaceLifIsis
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifIsisAuthentication(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceLifIsisBfdCfg(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceLifIsisCsnpIntervalList(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceLifIsisHelloIntervalList(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceLifIsisHelloIntervalMinimalList(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceLifIsisHelloMultiplierList(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceLifIsisMeshGroup(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceLifIsisMetricList(in["metric_list"].([]interface{}))
		ret.Network = in["network"].(string)
		ret.PasswordList = getSliceInterfaceLifIsisPasswordList(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceLifIsisPriorityList(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceLifIsisWideMetricList(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIsisAuthentication(d []interface{}) edpt.InterfaceLifIsisAuthentication {
	var ret edpt.InterfaceLifIsisAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceLifIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceLifIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceLifIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceLifIsisAuthenticationSendOnlyList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationSendOnlyList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceLifIsisAuthenticationModeList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationModeList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceLifIsisAuthenticationKeyChainList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationKeyChainList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsisBfdCfg(d []interface{}) edpt.InterfaceLifIsisBfdCfg {
	var ret edpt.InterfaceLifIsisBfdCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceLifIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceLifIsisCsnpIntervalList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisCsnpIntervalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloIntervalList(d []interface{}) []edpt.InterfaceLifIsisHelloIntervalList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloIntervalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceLifIsisHelloIntervalMinimalList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloIntervalMinimalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceLifIsisHelloMultiplierList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloMultiplierList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsisMeshGroup(d []interface{}) edpt.InterfaceLifIsisMeshGroup {
	var ret edpt.InterfaceLifIsisMeshGroup
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceLifIsisMetricList(d []interface{}) []edpt.InterfaceLifIsisMetricList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisMetricList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisPasswordList(d []interface{}) []edpt.InterfaceLifIsisPasswordList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisPasswordList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisPriorityList(d []interface{}) []edpt.InterfaceLifIsisPriorityList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisPriorityList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisWideMetricList(d []interface{}) []edpt.InterfaceLifIsisWideMetricList {
	count := len(d)
	ret := make([]edpt.InterfaceLifIsisWideMetricList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifSamplingEnable(d []interface{}) []edpt.InterfaceLifSamplingEnable {
	count := len(d)
	ret := make([]edpt.InterfaceLifSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
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
	ret.Inst.Bfd = getObjectInterfaceLifBfd(d.Get("bfd").([]interface{}))
	ret.Inst.Ifname = d.Get("ifname").(string)
	ret.Inst.Ip = getObjectInterfaceLifIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceLifIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceLifIsis(d.Get("isis").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SamplingEnable = getSliceInterfaceLifSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
