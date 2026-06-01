package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopback() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback`: Loopback interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackCreate,
		UpdateContext: resourceInterfaceLoopbackUpdate,
		ReadContext:   resourceInterfaceLoopbackRead,
		DeleteContext: resourceInterfaceLoopbackDelete,

		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Loopback interface number",
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
									},
									"ipv4_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
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
						"ospf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Name for the interface",
			},
			"snmp_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trap_source": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "The trap source",
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

func getObjectInterfaceLoopbackIp765(d []interface{}) edpt.InterfaceLoopbackIp765 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIp765
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceInterfaceLoopbackIpAddressList766(in["address_list"].([]interface{}))
		//omit uuid
		ret.Router = getObjectInterfaceLoopbackIpRouter767(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceLoopbackIpRip769(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpOspf777(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIpAddressList766(d []interface{}) []edpt.InterfaceLoopbackIpAddressList766 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpAddressList766, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpAddressList766
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRouter767(d []interface{}) edpt.InterfaceLoopbackIpRouter767 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRouter767
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceLoopbackIpRouterIsis768(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRouterIsis768(d []interface{}) edpt.InterfaceLoopbackIpRouterIsis768 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRouterIsis768
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpRip769(d []interface{}) edpt.InterfaceLoopbackIpRip769 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRip769
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLoopbackIpRipAuthentication770(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceLoopbackIpRipSendCfg774(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceLoopbackIpRipReceiveCfg775(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceLoopbackIpRipSplitHorizonCfg776(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthentication770(d []interface{}) edpt.InterfaceLoopbackIpRipAuthentication770 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthentication770
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceLoopbackIpRipAuthenticationStr771(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLoopbackIpRipAuthenticationMode772(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLoopbackIpRipAuthenticationKeyChain773(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationStr771(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationStr771 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationStr771
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationMode772(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationMode772 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationMode772
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationKeyChain773(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationKeyChain773 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationKeyChain773
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSendCfg774(d []interface{}) edpt.InterfaceLoopbackIpRipSendCfg774 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipSendCfg774
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipReceiveCfg775(d []interface{}) edpt.InterfaceLoopbackIpRipReceiveCfg775 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipReceiveCfg775
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSplitHorizonCfg776(d []interface{}) edpt.InterfaceLoopbackIpRipSplitHorizonCfg776 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipSplitHorizonCfg776
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspf777(d []interface{}) edpt.InterfaceLoopbackIpOspf777 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspf777
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceLoopbackIpOspfOspfGlobal778(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceLoopbackIpOspfOspfIpList784(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobal778(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobal778 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobal778
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg779(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalBfdCfg780(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg781(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg779(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg779 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg779
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalBfdCfg780(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalBfdCfg780 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalBfdCfg780
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg781(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg781 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg781
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5783(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5783(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5783 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5783
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfIpList784(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpList784 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpList784, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpList784
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6786(d []interface{}) edpt.InterfaceLoopbackIpv6786 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6786
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceInterfaceLoopbackIpv6AddressList787(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		//omit uuid
		ret.Router = getObjectInterfaceLoopbackIpv6Router788(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceLoopbackIpv6Rip793(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpv6Ospf795(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6AddressList787(d []interface{}) []edpt.InterfaceLoopbackIpv6AddressList787 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6AddressList787, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6AddressList787
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Anycast = in["anycast"].(int)
		oi.LinkLocal = in["link_local"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Router788(d []interface{}) edpt.InterfaceLoopbackIpv6Router788 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6Router788
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceLoopbackIpv6RouterRipng789(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpv6RouterOspf790(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceLoopbackIpv6RouterIsis792(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterRipng789(d []interface{}) edpt.InterfaceLoopbackIpv6RouterRipng789 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RouterRipng789
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterOspf790(d []interface{}) edpt.InterfaceLoopbackIpv6RouterOspf790 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RouterOspf790
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceLoopbackIpv6RouterOspfAreaList791(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6RouterOspfAreaList791(d []interface{}) []edpt.InterfaceLoopbackIpv6RouterOspfAreaList791 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6RouterOspfAreaList791, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6RouterOspfAreaList791
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterIsis792(d []interface{}) edpt.InterfaceLoopbackIpv6RouterIsis792 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RouterIsis792
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Rip793(d []interface{}) edpt.InterfaceLoopbackIpv6Rip793 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6Rip793
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceLoopbackIpv6RipSplitHorizonCfg794(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RipSplitHorizonCfg794(d []interface{}) edpt.InterfaceLoopbackIpv6RipSplitHorizonCfg794 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RipSplitHorizonCfg794
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Ospf795(d []interface{}) edpt.InterfaceLoopbackIpv6Ospf795 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6Ospf795
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceLoopbackIpv6OspfCostCfg796(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg797(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg798(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg799(in["mtu_ignore_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceLoopbackIpv6OspfPriorityCfg800(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg801(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg802(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfCostCfg796(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfCostCfg796 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfCostCfg796, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfCostCfg796
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg797(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg797 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg797, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg797
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg798(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg798 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg798, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg798
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg799(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg799 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg799, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg799
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfPriorityCfg800(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfPriorityCfg800 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfPriorityCfg800, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfPriorityCfg800
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg801(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg801 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg801, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg801
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg802(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg802 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg802, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg802
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsis803(d []interface{}) edpt.InterfaceLoopbackIsis803 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsis803
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLoopbackIsisAuthentication804(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceLoopbackIsisBfdCfg808(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceLoopbackIsisCsnpIntervalList809(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceLoopbackIsisHelloIntervalList810(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceLoopbackIsisHelloIntervalMinimalList811(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceLoopbackIsisHelloMultiplierList812(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceLoopbackIsisMeshGroup813(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceLoopbackIsisMetricList814(in["metric_list"].([]interface{}))
		ret.PasswordList = getSliceInterfaceLoopbackIsisPasswordList815(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceLoopbackIsisPriorityList816(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceLoopbackIsisWideMetricList817(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIsisAuthentication804(d []interface{}) edpt.InterfaceLoopbackIsisAuthentication804 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsisAuthentication804
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceLoopbackIsisAuthenticationSendOnlyList805(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceLoopbackIsisAuthenticationModeList806(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceLoopbackIsisAuthenticationKeyChainList807(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationSendOnlyList805(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList805 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList805, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList805
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationModeList806(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationModeList806 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationModeList806, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationModeList806
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationKeyChainList807(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationKeyChainList807 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationKeyChainList807, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationKeyChainList807
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsisBfdCfg808(d []interface{}) edpt.InterfaceLoopbackIsisBfdCfg808 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsisBfdCfg808
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisCsnpIntervalList809(d []interface{}) []edpt.InterfaceLoopbackIsisCsnpIntervalList809 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisCsnpIntervalList809, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisCsnpIntervalList809
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloIntervalList810(d []interface{}) []edpt.InterfaceLoopbackIsisHelloIntervalList810 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloIntervalList810, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloIntervalList810
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloIntervalMinimalList811(d []interface{}) []edpt.InterfaceLoopbackIsisHelloIntervalMinimalList811 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloIntervalMinimalList811, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloIntervalMinimalList811
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloMultiplierList812(d []interface{}) []edpt.InterfaceLoopbackIsisHelloMultiplierList812 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloMultiplierList812, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloMultiplierList812
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsisMeshGroup813(d []interface{}) edpt.InterfaceLoopbackIsisMeshGroup813 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsisMeshGroup813
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisMetricList814(d []interface{}) []edpt.InterfaceLoopbackIsisMetricList814 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisMetricList814, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisMetricList814
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisPasswordList815(d []interface{}) []edpt.InterfaceLoopbackIsisPasswordList815 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisPasswordList815, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisPasswordList815
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisPriorityList816(d []interface{}) []edpt.InterfaceLoopbackIsisPriorityList816 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisPriorityList816, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisPriorityList816
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisWideMetricList817(d []interface{}) []edpt.InterfaceLoopbackIsisWideMetricList817 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisWideMetricList817, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisWideMetricList817
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackSnmpServer(d []interface{}) edpt.InterfaceLoopbackSnmpServer {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackSnmpServer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TrapSource = in["trap_source"].(int)
	}
	return ret
}

func dataToEndpointInterfaceLoopback(d *schema.ResourceData) edpt.InterfaceLoopback {
	var ret edpt.InterfaceLoopback
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceLoopbackIp765(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceLoopbackIpv6786(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceLoopbackIsis803(d.Get("isis").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SnmpServer = getObjectInterfaceLoopbackSnmpServer(d.Get("snmp_server").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
