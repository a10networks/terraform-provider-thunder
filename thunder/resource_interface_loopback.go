package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInterfaceLoopback() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback`: Loopback interface\n\n",
		CreateContext: resourceInterfaceLoopbackCreate,
		UpdateContext: resourceInterfaceLoopbackUpdate,
		ReadContext:   resourceInterfaceLoopbackRead,
		DeleteContext: resourceInterfaceLoopbackDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Loopback interface number",
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Name for the interface",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"snmp_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trap_source": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "The trap source",
							ValidateFunc: validation.IntBetween(0, 1),
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

func resourceInterfaceLoopbackCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopback(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopback(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopback(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopback(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLoopbackIp(d []interface{}) edpt.InterfaceLoopbackIp {
	var ret edpt.InterfaceLoopbackIp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AddressList = getSliceInterfaceLoopbackIpAddressList(in["address_list"].([]interface{}))
		//omit uuid
		ret.Router = getObjectInterfaceLoopbackIpRouter(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceLoopbackIpRip(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpOspf(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIpAddressList(d []interface{}) []edpt.InterfaceLoopbackIpAddressList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpAddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRouter(d []interface{}) edpt.InterfaceLoopbackIpRouter {
	var ret edpt.InterfaceLoopbackIpRouter
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Isis = getObjectInterfaceLoopbackIpRouterIsis(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRouterIsis(d []interface{}) edpt.InterfaceLoopbackIpRouterIsis {
	var ret edpt.InterfaceLoopbackIpRouterIsis
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

func getObjectInterfaceLoopbackIpRip(d []interface{}) edpt.InterfaceLoopbackIpRip {
	var ret edpt.InterfaceLoopbackIpRip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceLoopbackIpRipAuthentication(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceLoopbackIpRipSendCfg(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceLoopbackIpRipReceiveCfg(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceLoopbackIpRipSplitHorizonCfg(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthentication(d []interface{}) edpt.InterfaceLoopbackIpRipAuthentication {
	var ret edpt.InterfaceLoopbackIpRipAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Str = getObjectInterfaceLoopbackIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLoopbackIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLoopbackIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationStr(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationStr {
	var ret edpt.InterfaceLoopbackIpRipAuthenticationStr
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationMode(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationMode {
	var ret edpt.InterfaceLoopbackIpRipAuthenticationMode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationKeyChain {
	var ret edpt.InterfaceLoopbackIpRipAuthenticationKeyChain
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSendCfg(d []interface{}) edpt.InterfaceLoopbackIpRipSendCfg {
	var ret edpt.InterfaceLoopbackIpRipSendCfg
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

func getObjectInterfaceLoopbackIpRipReceiveCfg(d []interface{}) edpt.InterfaceLoopbackIpRipReceiveCfg {
	var ret edpt.InterfaceLoopbackIpRipReceiveCfg
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

func getObjectInterfaceLoopbackIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceLoopbackIpRipSplitHorizonCfg {
	var ret edpt.InterfaceLoopbackIpRipSplitHorizonCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspf(d []interface{}) edpt.InterfaceLoopbackIpOspf {
	var ret edpt.InterfaceLoopbackIpOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceLoopbackIpOspfOspfGlobal(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceLoopbackIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobal(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobal {
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobal
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalBfdCfg(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg {
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg
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

func getObjectInterfaceLoopbackIpOspfOspfGlobalBfdCfg(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalBfdCfg {
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalBfdCfg
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

func getObjectInterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg {
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg
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

func getSliceInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5 {
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5
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

func getSliceInterfaceLoopbackIpOspfOspfIpList(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6(d []interface{}) edpt.InterfaceLoopbackIpv6 {
	var ret edpt.InterfaceLoopbackIpv6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AddressList = getSliceInterfaceLoopbackIpv6AddressList(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceLoopbackIpv6Router(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceLoopbackIpv6Rip(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpv6Ospf(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6AddressList(d []interface{}) []edpt.InterfaceLoopbackIpv6AddressList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6AddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Anycast = in["anycast"].(int)
		oi.LinkLocal = in["link_local"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Router(d []interface{}) edpt.InterfaceLoopbackIpv6Router {
	var ret edpt.InterfaceLoopbackIpv6Router
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ripng = getObjectInterfaceLoopbackIpv6RouterRipng(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpv6RouterOspf(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceLoopbackIpv6RouterIsis(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterRipng(d []interface{}) edpt.InterfaceLoopbackIpv6RouterRipng {
	var ret edpt.InterfaceLoopbackIpv6RouterRipng
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

func getObjectInterfaceLoopbackIpv6RouterOspf(d []interface{}) edpt.InterfaceLoopbackIpv6RouterOspf {
	var ret edpt.InterfaceLoopbackIpv6RouterOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AreaList = getSliceInterfaceLoopbackIpv6RouterOspfAreaList(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6RouterOspfAreaList(d []interface{}) []edpt.InterfaceLoopbackIpv6RouterOspfAreaList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6RouterOspfAreaList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6RouterOspfAreaList
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterIsis(d []interface{}) edpt.InterfaceLoopbackIpv6RouterIsis {
	var ret edpt.InterfaceLoopbackIpv6RouterIsis
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

func getObjectInterfaceLoopbackIpv6Rip(d []interface{}) edpt.InterfaceLoopbackIpv6Rip {
	var ret edpt.InterfaceLoopbackIpv6Rip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceLoopbackIpv6RipSplitHorizonCfg(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RipSplitHorizonCfg(d []interface{}) edpt.InterfaceLoopbackIpv6RipSplitHorizonCfg {
	var ret edpt.InterfaceLoopbackIpv6RipSplitHorizonCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Ospf(d []interface{}) edpt.InterfaceLoopbackIpv6Ospf {
	var ret edpt.InterfaceLoopbackIpv6Ospf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceLoopbackIpv6OspfCostCfg(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg(in["mtu_ignore_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceLoopbackIpv6OspfPriorityCfg(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfCostCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfCostCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfPriorityCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfPriorityCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsis(d []interface{}) edpt.InterfaceLoopbackIsis {
	var ret edpt.InterfaceLoopbackIsis
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceLoopbackIsisAuthentication(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceLoopbackIsisBfdCfg(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceLoopbackIsisCsnpIntervalList(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceLoopbackIsisHelloIntervalList(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceLoopbackIsisHelloIntervalMinimalList(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceLoopbackIsisHelloMultiplierList(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceLoopbackIsisMeshGroup(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceLoopbackIsisMetricList(in["metric_list"].([]interface{}))
		ret.PasswordList = getSliceInterfaceLoopbackIsisPasswordList(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceLoopbackIsisPriorityList(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceLoopbackIsisWideMetricList(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIsisAuthentication(d []interface{}) edpt.InterfaceLoopbackIsisAuthentication {
	var ret edpt.InterfaceLoopbackIsisAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceLoopbackIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceLoopbackIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceLoopbackIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationModeList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationModeList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationKeyChainList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationKeyChainList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsisBfdCfg(d []interface{}) edpt.InterfaceLoopbackIsisBfdCfg {
	var ret edpt.InterfaceLoopbackIsisBfdCfg
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

func getSliceInterfaceLoopbackIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceLoopbackIsisCsnpIntervalList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisCsnpIntervalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloIntervalList(d []interface{}) []edpt.InterfaceLoopbackIsisHelloIntervalList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloIntervalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceLoopbackIsisHelloIntervalMinimalList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloIntervalMinimalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceLoopbackIsisHelloMultiplierList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloMultiplierList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsisMeshGroup(d []interface{}) edpt.InterfaceLoopbackIsisMeshGroup {
	var ret edpt.InterfaceLoopbackIsisMeshGroup
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

func getSliceInterfaceLoopbackIsisMetricList(d []interface{}) []edpt.InterfaceLoopbackIsisMetricList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisMetricList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisPasswordList(d []interface{}) []edpt.InterfaceLoopbackIsisPasswordList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisPasswordList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisPriorityList(d []interface{}) []edpt.InterfaceLoopbackIsisPriorityList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisPriorityList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisWideMetricList(d []interface{}) []edpt.InterfaceLoopbackIsisWideMetricList {
	count := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisWideMetricList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackSnmpServer(d []interface{}) edpt.InterfaceLoopbackSnmpServer {
	var ret edpt.InterfaceLoopbackSnmpServer
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.TrapSource = in["trap_source"].(int)
	}
	return ret
}

func dataToEndpointInterfaceLoopback(d *schema.ResourceData) edpt.InterfaceLoopback {
	var ret edpt.InterfaceLoopback
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceLoopbackIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceLoopbackIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceLoopbackIsis(d.Get("isis").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SnmpServer = getObjectInterfaceLoopbackSnmpServer(d.Get("snmp_server").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
