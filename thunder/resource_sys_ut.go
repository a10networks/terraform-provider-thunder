package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut`: System unit test\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtCreate,
		UpdateContext: resourceSysUtUpdate,
		ReadContext:   resourceSysUtRead,
		DeleteContext: resourceSysUtDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Set device to SUT mode; 'disable': Set device to normal mode;",
			},
			"common": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proceed_on_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run test even in case of event failure",
						},
						"delay": {
							Type: schema.TypeInt, Optional: true, Description: "wait time in seconds after each run",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"event_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event_number": {
							Type: schema.TypeInt, Required: true, Description: "Event number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"action_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": {
										Type: schema.TypeString, Required: true, Description: "'send': Test event; 'expect': Expected result; 'wait': Introduce a delay;",
									},
									"template": {
										Type: schema.TypeString, Optional: true, Description: "Packet template",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
									},
									"delay": {
										Type: schema.TypeInt, Optional: true, Description: "Delay in seconds",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"l1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"eth_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ethernet_start": {
																Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
															},
															"ethernet_end": {
																Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
															},
														},
													},
												},
												"trunk_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trunk_start": {
																Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
															},
															"trunk_end": {
																Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
															},
														},
													},
												},
												"length": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet length",
												},
												"value": {
													Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at L2 header",
												},
												"auto": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"l2": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethertype": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "L2 frame type",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'arp': arp; 'ipv4': ipv4; 'ipv6': ipv6;",
												},
												"value": {
													Type: schema.TypeInt, Optional: true, Description: "ethertype number",
												},
												"vlan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Vlan ID on the packet. 0 is untagged",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"mac_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"src_dst": {
																Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
															},
															"address_type": {
																Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
															},
															"virtual_server": {
																Type: schema.TypeString, Optional: true, Description: "vip",
															},
															"nat_pool": {
																Type: schema.TypeString, Optional: true, Description: "Nat pool",
															},
															"ethernet": {
																Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
															},
															"ve": {
																Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
															},
															"trunk": {
																Type: schema.TypeInt, Optional: true, Description: "Trunk number",
															},
															"value": {
																Type: schema.TypeString, Optional: true, Description: "Mac Address",
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
									"l3": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "L4 Protocol",
												},
												"type": {
													Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp;",
												},
												"value": {
													Type: schema.TypeInt, Optional: true, Description: "protocol number",
												},
												"checksum": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"ttl": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"ip_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"src_dst": {
																Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
															},
															"ipv4_address": {
																Type: schema.TypeString, Optional: true, Description: "IP address",
															},
															"ipv6_address": {
																Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
															},
															"nat_pool": {
																Type: schema.TypeString, Optional: true, Description: "Nat pool",
															},
															"virtual_server": {
																Type: schema.TypeString, Optional: true, Description: "vip",
															},
															"ethernet": {
																Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
															},
															"ve": {
																Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
															},
															"trunk": {
																Type: schema.TypeInt, Optional: true, Description: "Trunk number",
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
									"tcp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_port": {
													Type: schema.TypeInt, Optional: true, Description: "Source port value",
												},
												"dest_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
												},
												"dest_port_value": {
													Type: schema.TypeInt, Optional: true, Description: "Dest port value",
												},
												"nat_pool": {
													Type: schema.TypeString, Optional: true, Description: "Nat pool port",
												},
												"seq_number": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"ack_seq_number": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"checksum": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"urgent": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"window": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"flags": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"syn": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
															},
															"ack": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
															},
															"fin": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
															},
															"rst": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
															},
															"psh": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
															},
															"ece": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
															},
															"urg": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
															},
															"cwr": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"options": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mss": {
																Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
															},
															"wscale": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"sack_type": {
																Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
															},
															"time_stamp_enable": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
															},
															"nop": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
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
									"udp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_port": {
													Type: schema.TypeInt, Optional: true, Description: "Source port value",
												},
												"dest_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
												},
												"dest_port_value": {
													Type: schema.TypeInt, Optional: true, Description: "Dest port value",
												},
												"nat_pool": {
													Type: schema.TypeString, Optional: true, Description: "Nat pool port",
												},
												"length": {
													Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at UDP header",
												},
												"checksum": {
													Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"ignore_validation": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"l1": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate TX descriptor. This includes Tx port, Len & vlan",
												},
												"l2": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L2 header",
												},
												"l3": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L3 header",
												},
												"l4": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L4 header",
												},
												"all": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skip validation",
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
			"run_test": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "'basic': Run Basic mode; 'fault-injection': Run FI mode. This will also run Basic mode to gather data; 'cpu-rr': Run CPU RR mode; 'frag': Run IP frag mode;",
						},
					},
				},
			},
			"secondary_name": {
				Type: schema.TypeString, Optional: true, Description: "Minor test case name. This is autogenerated",
			},
			"state_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Node name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"next_state_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Node name",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"case_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"case_number": {
													Type: schema.TypeInt, Required: true, Description: "",
												},
												"repeat": {
													Type: schema.TypeInt, Optional: true, Description: "number of times case should be repeated before state transition",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
												},
												"action_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"direction": {
																Type: schema.TypeString, Required: true, Description: "'send': Test event; 'expect': Expected result; 'wait': Introduce a delay;",
															},
															"template": {
																Type: schema.TypeString, Optional: true, Description: "Packet template",
															},
															"drop": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
															},
															"delay": {
																Type: schema.TypeInt, Optional: true, Description: "Delay in seconds",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"l1": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"eth_list": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"ethernet_start": {
																						Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
																					},
																					"ethernet_end": {
																						Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
																					},
																				},
																			},
																		},
																		"trunk_list": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"trunk_start": {
																						Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
																					},
																					"trunk_end": {
																						Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
																					},
																				},
																			},
																		},
																		"length": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet length",
																		},
																		"value": {
																			Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at L2 header",
																		},
																		"auto": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																	},
																},
															},
															"l2": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ethertype": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "L2 frame type",
																		},
																		"protocol": {
																			Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'arp': arp; 'ipv4': ipv4; 'ipv6': ipv6;",
																		},
																		"value": {
																			Type: schema.TypeInt, Optional: true, Description: "ethertype number",
																		},
																		"vlan": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Vlan ID on the packet. 0 is untagged",
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"mac_list": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"src_dst": {
																						Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
																					},
																					"address_type": {
																						Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
																					},
																					"virtual_server": {
																						Type: schema.TypeString, Optional: true, Description: "vip",
																					},
																					"nat_pool": {
																						Type: schema.TypeString, Optional: true, Description: "Nat pool",
																					},
																					"ethernet": {
																						Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
																					},
																					"ve": {
																						Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
																					},
																					"trunk": {
																						Type: schema.TypeInt, Optional: true, Description: "Trunk number",
																					},
																					"value": {
																						Type: schema.TypeString, Optional: true, Description: "Mac Address",
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
															"l3": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"protocol": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "L4 Protocol",
																		},
																		"type": {
																			Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp;",
																		},
																		"value": {
																			Type: schema.TypeInt, Optional: true, Description: "protocol number",
																		},
																		"checksum": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
																		},
																		"ttl": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"ip_list": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"src_dst": {
																						Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
																					},
																					"ipv4_address": {
																						Type: schema.TypeString, Optional: true, Description: "IP address",
																					},
																					"ipv6_address": {
																						Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
																					},
																					"virtual_server": {
																						Type: schema.TypeString, Optional: true, Description: "vip",
																					},
																					"nat_pool": {
																						Type: schema.TypeString, Optional: true, Description: "Nat pool",
																					},
																					"ethernet": {
																						Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
																					},
																					"ve": {
																						Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
																					},
																					"trunk": {
																						Type: schema.TypeInt, Optional: true, Description: "Trunk number",
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
															"tcp": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"src_port": {
																			Type: schema.TypeInt, Optional: true, Description: "Source port value",
																		},
																		"dest_port": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
																		},
																		"dest_port_value": {
																			Type: schema.TypeInt, Optional: true, Description: "Dest port value",
																		},
																		"nat_pool": {
																			Type: schema.TypeString, Optional: true, Description: "Nat pool port",
																		},
																		"seq_number": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
																		},
																		"ack_seq_number": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
																		},
																		"checksum": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
																		},
																		"urgent": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
																		},
																		"window": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																		},
																		"flags": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"syn": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
																					},
																					"ack": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
																					},
																					"fin": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
																					},
																					"rst": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
																					},
																					"psh": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
																					},
																					"ece": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
																					},
																					"urg": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
																					},
																					"cwr": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
																					},
																					"uuid": {
																						Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
																					},
																				},
																			},
																		},
																		"options": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"mss": {
																						Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
																					},
																					"wscale": {
																						Type: schema.TypeInt, Optional: true, Description: "",
																					},
																					"sack_type": {
																						Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
																					},
																					"time_stamp_enable": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
																					},
																					"nop": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
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
															"udp": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"src_port": {
																			Type: schema.TypeInt, Optional: true, Description: "Source port value",
																		},
																		"dest_port": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
																		},
																		"dest_port_value": {
																			Type: schema.TypeInt, Optional: true, Description: "Dest port value",
																		},
																		"nat_pool": {
																			Type: schema.TypeString, Optional: true, Description: "Nat pool port",
																		},
																		"length": {
																			Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at UDP header",
																		},
																		"checksum": {
																			Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
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
								},
							},
						},
					},
				},
			},
			"template_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "template name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"ignore_validation": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"l1": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate TX descriptor. This includes Tx port, Len & vlan",
									},
									"l2": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L2 header",
									},
									"l3": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L3 header",
									},
									"l4": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L4 header",
									},
									"all": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skip validation",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"l1": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eth_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet_start": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
												},
												"ethernet_end": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
												},
											},
										},
									},
									"trunk_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trunk_start": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
												},
												"trunk_end": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
												},
											},
										},
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
									},
									"length": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet length",
									},
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at L2 header",
									},
									"auto": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"l2": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethertype": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "L2 frame type",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'arp': arp; 'ipv4': ipv4; 'ipv6': ipv6;",
									},
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "ethertype number",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Vlan ID on the packet. 0 is untagged",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"mac_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_dst": {
													Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
												},
												"address_type": {
													Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
												},
												"virtual_server": {
													Type: schema.TypeString, Optional: true, Description: "vip",
												},
												"nat_pool": {
													Type: schema.TypeString, Optional: true, Description: "Nat pool",
												},
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
												},
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
												},
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk number",
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "Mac Address",
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
						"l3": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "L4 Protocol",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp;",
									},
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "protocol number",
									},
									"checksum": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"ip_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_dst": {
													Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
												},
												"ipv4_start_address": {
													Type: schema.TypeString, Optional: true, Description: "IP address",
												},
												"ipv4_end_address": {
													Type: schema.TypeString, Optional: true, Description: "IP end address",
												},
												"ipv6_start_address": {
													Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
												},
												"ipv6_end_address": {
													Type: schema.TypeString, Optional: true, Description: "Ipv6 end address",
												},
												"virtual_server": {
													Type: schema.TypeString, Optional: true, Description: "vip",
												},
												"nat_pool": {
													Type: schema.TypeString, Optional: true, Description: "Nat pool",
												},
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
												},
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
												},
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk number",
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
						"tcp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_port_range": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "Source port value",
												},
												"src_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "Src port end value",
												},
											},
										},
									},
									"dest_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
									},
									"dest_port_value": {
										Type: schema.TypeInt, Optional: true, Description: "Dest port value",
									},
									"nat_pool": {
										Type: schema.TypeString, Optional: true, Description: "Nat pool port",
									},
									"seq_number": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"ack_seq_number": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"checksum": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"urgent": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"window": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"flags": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"syn": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
												},
												"ack": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
												},
												"fin": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
												},
												"rst": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
												},
												"psh": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
												},
												"ece": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
												},
												"urg": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
												},
												"cwr": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"options": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mss": {
													Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
												},
												"wscale": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sack_type": {
													Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
												},
												"time_stamp_enable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
												},
												"nop": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
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
						"udp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_port_range": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "Source port value",
												},
												"src_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "Src port end value",
												},
											},
										},
									},
									"dest_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
									},
									"dest_port_value": {
										Type: schema.TypeInt, Optional: true, Description: "Dest port value",
									},
									"nat_pool": {
										Type: schema.TypeString, Optional: true, Description: "Nat pool port",
									},
									"length": {
										Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at UDP header",
									},
									"checksum": {
										Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
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
			"test_name": {
				Type: schema.TypeString, Optional: true, Description: "Name for this test case",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSysUtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSysUtCommon1569(d []interface{}) edpt.SysUtCommon1569 {

	count1 := len(d)
	var ret edpt.SysUtCommon1569
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProceedOnError = in["proceed_on_error"].(int)
		ret.Delay = in["delay"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtEventList(d []interface{}) []edpt.SysUtEventList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventList
		oi.EventNumber = in["event_number"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.ActionList = getSliceSysUtEventListActionList(in["action_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtEventListActionList(d []interface{}) []edpt.SysUtEventListActionList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventListActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventListActionList
		oi.Direction = in["direction"].(string)
		oi.Template = in["template"].(string)
		oi.Drop = in["drop"].(int)
		oi.Delay = in["delay"].(int)
		//omit uuid
		oi.L1 = getObjectSysUtEventListActionListL1(in["l1"].([]interface{}))
		oi.L2 = getObjectSysUtEventListActionListL2(in["l2"].([]interface{}))
		oi.L3 = getObjectSysUtEventListActionListL3(in["l3"].([]interface{}))
		oi.Tcp = getObjectSysUtEventListActionListTcp(in["tcp"].([]interface{}))
		oi.Udp = getObjectSysUtEventListActionListUdp(in["udp"].([]interface{}))
		oi.IgnoreValidation = getObjectSysUtEventListActionListIgnoreValidation(in["ignore_validation"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtEventListActionListL1(d []interface{}) edpt.SysUtEventListActionListL1 {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListL1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthList = getSliceSysUtEventListActionListL1EthList(in["eth_list"].([]interface{}))
		ret.Trunk_list = getSliceSysUtEventListActionListL1Trunk_list(in["trunk_list"].([]interface{}))
		ret.Length = in["length"].(int)
		ret.Value = in["value"].(int)
		ret.Auto = in["auto"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtEventListActionListL1EthList(d []interface{}) []edpt.SysUtEventListActionListL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventListActionListL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventListActionListL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtEventListActionListL1Trunk_list(d []interface{}) []edpt.SysUtEventListActionListL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtEventListActionListL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventListActionListL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtEventListActionListL2(d []interface{}) edpt.SysUtEventListActionListL2 {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListL2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethertype = in["ethertype"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Value = in["value"].(int)
		ret.Vlan = in["vlan"].(int)
		//omit uuid
		ret.MacList = getSliceSysUtEventListActionListL2MacList(in["mac_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtEventListActionListL2MacList(d []interface{}) []edpt.SysUtEventListActionListL2MacList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventListActionListL2MacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventListActionListL2MacList
		oi.SrcDst = in["src_dst"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Value = in["value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtEventListActionListL3(d []interface{}) edpt.SysUtEventListActionListL3 {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListL3
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(int)
		ret.Type = in["type"].(string)
		ret.Value = in["value"].(int)
		ret.Checksum = in["checksum"].(string)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.IpList = getSliceSysUtEventListActionListL3IpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtEventListActionListL3IpList(d []interface{}) []edpt.SysUtEventListActionListL3IpList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventListActionListL3IpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventListActionListL3IpList
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtEventListActionListTcp(d []interface{}) edpt.SysUtEventListActionListTcp {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.SeqNumber = in["seq_number"].(string)
		ret.AckSeqNumber = in["ack_seq_number"].(string)
		ret.Checksum = in["checksum"].(string)
		ret.Urgent = in["urgent"].(string)
		ret.Window = in["window"].(string)
		//omit uuid
		ret.Flags = getObjectSysUtEventListActionListTcpFlags(in["flags"].([]interface{}))
		ret.Options = getObjectSysUtEventListActionListTcpOptions(in["options"].([]interface{}))
	}
	return ret
}

func getObjectSysUtEventListActionListTcpFlags(d []interface{}) edpt.SysUtEventListActionListTcpFlags {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListTcpFlags
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn = in["syn"].(int)
		ret.Ack = in["ack"].(int)
		ret.Fin = in["fin"].(int)
		ret.Rst = in["rst"].(int)
		ret.Psh = in["psh"].(int)
		ret.Ece = in["ece"].(int)
		ret.Urg = in["urg"].(int)
		ret.Cwr = in["cwr"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtEventListActionListTcpOptions(d []interface{}) edpt.SysUtEventListActionListTcpOptions {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListTcpOptions
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mss = in["mss"].(int)
		ret.Wscale = in["wscale"].(int)
		ret.SackType = in["sack_type"].(string)
		ret.TimeStampEnable = in["time_stamp_enable"].(int)
		ret.Nop = in["nop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtEventListActionListUdp(d []interface{}) edpt.SysUtEventListActionListUdp {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListUdp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.Length = in["length"].(int)
		ret.Checksum = in["checksum"].(string)
		//omit uuid
	}
	return ret
}

func getObjectSysUtEventListActionListIgnoreValidation(d []interface{}) edpt.SysUtEventListActionListIgnoreValidation {

	count1 := len(d)
	var ret edpt.SysUtEventListActionListIgnoreValidation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L1 = in["l1"].(int)
		ret.L2 = in["l2"].(int)
		ret.L3 = in["l3"].(int)
		ret.L4 = in["l4"].(int)
		ret.All = in["all"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtRunTest1570(d []interface{}) edpt.SysUtRunTest1570 {

	count1 := len(d)
	var ret edpt.SysUtRunTest1570
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getSliceSysUtStateList(d []interface{}) []edpt.SysUtStateList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.NextStateList = getSliceSysUtStateListNextStateList(in["next_state_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateListNextStateList(d []interface{}) []edpt.SysUtStateListNextStateList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.CaseList = getSliceSysUtStateListNextStateListCaseList(in["case_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateListNextStateListCaseList(d []interface{}) []edpt.SysUtStateListNextStateListCaseList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateListCaseList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateListCaseList
		oi.CaseNumber = in["case_number"].(int)
		oi.Repeat = in["repeat"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.ActionList = getSliceSysUtStateListNextStateListCaseListActionList(in["action_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateListNextStateListCaseListActionList(d []interface{}) []edpt.SysUtStateListNextStateListCaseListActionList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateListCaseListActionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateListCaseListActionList
		oi.Direction = in["direction"].(string)
		oi.Template = in["template"].(string)
		oi.Drop = in["drop"].(int)
		oi.Delay = in["delay"].(int)
		//omit uuid
		oi.L1 = getObjectSysUtStateListNextStateListCaseListActionListL1(in["l1"].([]interface{}))
		oi.L2 = getObjectSysUtStateListNextStateListCaseListActionListL2(in["l2"].([]interface{}))
		oi.L3 = getObjectSysUtStateListNextStateListCaseListActionListL3(in["l3"].([]interface{}))
		oi.Tcp = getObjectSysUtStateListNextStateListCaseListActionListTcp(in["tcp"].([]interface{}))
		oi.Udp = getObjectSysUtStateListNextStateListCaseListActionListUdp(in["udp"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListL1(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListL1 {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListL1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthList = getSliceSysUtStateListNextStateListCaseListActionListL1EthList(in["eth_list"].([]interface{}))
		ret.Trunk_list = getSliceSysUtStateListNextStateListCaseListActionListL1Trunk_list(in["trunk_list"].([]interface{}))
		ret.Length = in["length"].(int)
		ret.Value = in["value"].(int)
		ret.Auto = in["auto"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtStateListNextStateListCaseListActionListL1EthList(d []interface{}) []edpt.SysUtStateListNextStateListCaseListActionListL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateListCaseListActionListL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateListCaseListActionListL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateListNextStateListCaseListActionListL1Trunk_list(d []interface{}) []edpt.SysUtStateListNextStateListCaseListActionListL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateListCaseListActionListL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateListCaseListActionListL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListL2(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListL2 {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListL2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethertype = in["ethertype"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Value = in["value"].(int)
		ret.Vlan = in["vlan"].(int)
		//omit uuid
		ret.MacList = getSliceSysUtStateListNextStateListCaseListActionListL2MacList(in["mac_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtStateListNextStateListCaseListActionListL2MacList(d []interface{}) []edpt.SysUtStateListNextStateListCaseListActionListL2MacList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateListCaseListActionListL2MacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateListCaseListActionListL2MacList
		oi.SrcDst = in["src_dst"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Value = in["value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListL3(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListL3 {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListL3
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(int)
		ret.Type = in["type"].(string)
		ret.Value = in["value"].(int)
		ret.Checksum = in["checksum"].(string)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.IpList = getSliceSysUtStateListNextStateListCaseListActionListL3IpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtStateListNextStateListCaseListActionListL3IpList(d []interface{}) []edpt.SysUtStateListNextStateListCaseListActionListL3IpList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateListNextStateListCaseListActionListL3IpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateListNextStateListCaseListActionListL3IpList
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListTcp(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListTcp {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.SeqNumber = in["seq_number"].(string)
		ret.AckSeqNumber = in["ack_seq_number"].(string)
		ret.Checksum = in["checksum"].(string)
		ret.Urgent = in["urgent"].(string)
		ret.Window = in["window"].(string)
		//omit uuid
		ret.Flags = getObjectSysUtStateListNextStateListCaseListActionListTcpFlags(in["flags"].([]interface{}))
		ret.Options = getObjectSysUtStateListNextStateListCaseListActionListTcpOptions(in["options"].([]interface{}))
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListTcpFlags(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListTcpFlags {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListTcpFlags
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn = in["syn"].(int)
		ret.Ack = in["ack"].(int)
		ret.Fin = in["fin"].(int)
		ret.Rst = in["rst"].(int)
		ret.Psh = in["psh"].(int)
		ret.Ece = in["ece"].(int)
		ret.Urg = in["urg"].(int)
		ret.Cwr = in["cwr"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListTcpOptions(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListTcpOptions {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListTcpOptions
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mss = in["mss"].(int)
		ret.Wscale = in["wscale"].(int)
		ret.SackType = in["sack_type"].(string)
		ret.TimeStampEnable = in["time_stamp_enable"].(int)
		ret.Nop = in["nop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtStateListNextStateListCaseListActionListUdp(d []interface{}) edpt.SysUtStateListNextStateListCaseListActionListUdp {

	count1 := len(d)
	var ret edpt.SysUtStateListNextStateListCaseListActionListUdp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.Length = in["length"].(int)
		ret.Checksum = in["checksum"].(string)
		//omit uuid
	}
	return ret
}

func getSliceSysUtTemplateList(d []interface{}) []edpt.SysUtTemplateList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.IgnoreValidation = getObjectSysUtTemplateListIgnoreValidation(in["ignore_validation"].([]interface{}))
		oi.L1 = getObjectSysUtTemplateListL1(in["l1"].([]interface{}))
		oi.L2 = getObjectSysUtTemplateListL2(in["l2"].([]interface{}))
		oi.L3 = getObjectSysUtTemplateListL3(in["l3"].([]interface{}))
		oi.Tcp = getObjectSysUtTemplateListTcp(in["tcp"].([]interface{}))
		oi.Udp = getObjectSysUtTemplateListUdp(in["udp"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateListIgnoreValidation(d []interface{}) edpt.SysUtTemplateListIgnoreValidation {

	count1 := len(d)
	var ret edpt.SysUtTemplateListIgnoreValidation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L1 = in["l1"].(int)
		ret.L2 = in["l2"].(int)
		ret.L3 = in["l3"].(int)
		ret.L4 = in["l4"].(int)
		ret.All = in["all"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtTemplateListL1(d []interface{}) edpt.SysUtTemplateListL1 {

	count1 := len(d)
	var ret edpt.SysUtTemplateListL1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthList = getSliceSysUtTemplateListL1EthList(in["eth_list"].([]interface{}))
		ret.Trunk_list = getSliceSysUtTemplateListL1Trunk_list(in["trunk_list"].([]interface{}))
		ret.Drop = in["drop"].(int)
		ret.Length = in["length"].(int)
		ret.Value = in["value"].(int)
		ret.Auto = in["auto"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtTemplateListL1EthList(d []interface{}) []edpt.SysUtTemplateListL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateListL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateListL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtTemplateListL1Trunk_list(d []interface{}) []edpt.SysUtTemplateListL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateListL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateListL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateListL2(d []interface{}) edpt.SysUtTemplateListL2 {

	count1 := len(d)
	var ret edpt.SysUtTemplateListL2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethertype = in["ethertype"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Value = in["value"].(int)
		ret.Vlan = in["vlan"].(int)
		//omit uuid
		ret.MacList = getSliceSysUtTemplateListL2MacList(in["mac_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtTemplateListL2MacList(d []interface{}) []edpt.SysUtTemplateListL2MacList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateListL2MacList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateListL2MacList
		oi.SrcDst = in["src_dst"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Value = in["value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateListL3(d []interface{}) edpt.SysUtTemplateListL3 {

	count1 := len(d)
	var ret edpt.SysUtTemplateListL3
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(int)
		ret.Type = in["type"].(string)
		ret.Value = in["value"].(int)
		ret.Checksum = in["checksum"].(string)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.IpList = getSliceSysUtTemplateListL3IpList(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtTemplateListL3IpList(d []interface{}) []edpt.SysUtTemplateListL3IpList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateListL3IpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateListL3IpList
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4StartAddress = in["ipv4_start_address"].(string)
		oi.Ipv4EndAddress = in["ipv4_end_address"].(string)
		oi.Ipv6StartAddress = in["ipv6_start_address"].(string)
		oi.Ipv6EndAddress = in["ipv6_end_address"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateListTcp(d []interface{}) edpt.SysUtTemplateListTcp {

	count1 := len(d)
	var ret edpt.SysUtTemplateListTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPortRange = getSliceSysUtTemplateListTcpSrcPortRange(in["src_port_range"].([]interface{}))
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.SeqNumber = in["seq_number"].(string)
		ret.AckSeqNumber = in["ack_seq_number"].(string)
		ret.Checksum = in["checksum"].(string)
		ret.Urgent = in["urgent"].(string)
		ret.Window = in["window"].(string)
		//omit uuid
		ret.Flags = getObjectSysUtTemplateListTcpFlags(in["flags"].([]interface{}))
		ret.Options = getObjectSysUtTemplateListTcpOptions(in["options"].([]interface{}))
	}
	return ret
}

func getSliceSysUtTemplateListTcpSrcPortRange(d []interface{}) []edpt.SysUtTemplateListTcpSrcPortRange {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateListTcpSrcPortRange, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateListTcpSrcPortRange
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateListTcpFlags(d []interface{}) edpt.SysUtTemplateListTcpFlags {

	count1 := len(d)
	var ret edpt.SysUtTemplateListTcpFlags
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn = in["syn"].(int)
		ret.Ack = in["ack"].(int)
		ret.Fin = in["fin"].(int)
		ret.Rst = in["rst"].(int)
		ret.Psh = in["psh"].(int)
		ret.Ece = in["ece"].(int)
		ret.Urg = in["urg"].(int)
		ret.Cwr = in["cwr"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtTemplateListTcpOptions(d []interface{}) edpt.SysUtTemplateListTcpOptions {

	count1 := len(d)
	var ret edpt.SysUtTemplateListTcpOptions
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mss = in["mss"].(int)
		ret.Wscale = in["wscale"].(int)
		ret.SackType = in["sack_type"].(string)
		ret.TimeStampEnable = in["time_stamp_enable"].(int)
		ret.Nop = in["nop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtTemplateListUdp(d []interface{}) edpt.SysUtTemplateListUdp {

	count1 := len(d)
	var ret edpt.SysUtTemplateListUdp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPortRange = getSliceSysUtTemplateListUdpSrcPortRange(in["src_port_range"].([]interface{}))
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.Length = in["length"].(int)
		ret.Checksum = in["checksum"].(string)
		//omit uuid
	}
	return ret
}

func getSliceSysUtTemplateListUdpSrcPortRange(d []interface{}) []edpt.SysUtTemplateListUdpSrcPortRange {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateListUdpSrcPortRange, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateListUdpSrcPortRange
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUt(d *schema.ResourceData) edpt.SysUt {
	var ret edpt.SysUt
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Common = getObjectSysUtCommon1569(d.Get("common").([]interface{}))
	ret.Inst.EventList = getSliceSysUtEventList(d.Get("event_list").([]interface{}))
	ret.Inst.RunTest = getObjectSysUtRunTest1570(d.Get("run_test").([]interface{}))
	ret.Inst.SecondaryName = d.Get("secondary_name").(string)
	ret.Inst.StateList = getSliceSysUtStateList(d.Get("state_list").([]interface{}))
	ret.Inst.TemplateList = getSliceSysUtTemplateList(d.Get("template_list").([]interface{}))
	ret.Inst.TestName = d.Get("test_name").(string)
	//omit uuid
	return ret
}
