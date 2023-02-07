package thunder

import (
	"context"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceInterfaceVe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve`: Virtual ethernet interface\n\n",
		CreateContext: resourceInterfaceVeCreate,
		UpdateContext: resourceInterfaceVeUpdate,
		ReadContext:   resourceInterfaceVeRead,
		DeleteContext: resourceInterfaceVeDelete,
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
							Type: schema.TypeString, Optional: true, Description: "Named Access List",
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
			"ddos": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DDoS inside (trusted) or outside (untrusted) interface",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DDoS inside (trusted) or outside (untrusted) interface",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"icmp_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"normal": {
							Type: schema.TypeInt, Optional: true, Description: "Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"lockup": {
							Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"lockup_period": {
							Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
							ValidateFunc: validation.IntBetween(1, 16383),
						},
					},
				},
			},
			"icmpv6_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"normal_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"lockup_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"lockup_period_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
							ValidateFunc: validation.IntBetween(1, 16383),
						},
					},
				},
			},
			"ifnum": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Virtual ethernet interface number",
				ValidateFunc: validation.IntBetween(2, 4094),
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
						"client": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client facing interface for IPv4/v6 traffic",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server facing interface for IPv4/v6 traffic",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"helper_address_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"helper_address": {
										Type: schema.TypeString, Optional: true, Description: "Helper address for DHCP packets (IP address)",
										ValidateFunc: validation.IsIPv4Address,
									},
								},
							},
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"ttl_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"syn_cookie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SYN-cookie on the interface",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"slb_partition_redirect": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Redirect SLB traffic across partition",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"stateful_firewall": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inside (private) interface for stateful firewall",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"outside": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outside (public) interface for stateful firewall",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"access_list": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Access-list for traffic from the outside",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"acl_id": {
										Type: schema.TypeInt, Optional: true, Description: "ACL id",
										ValidateFunc: validation.IntBetween(1, 199),
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
																ValidateFunc:  validation.IntBetween(0, 1),
															},
															"non_broadcast": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF NBMA network",
																ValidateFunc:  validation.IntBetween(0, 1),
															},
															"point_to_point": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-point network",
																ValidateFunc:  validation.IntBetween(0, 1),
															},
															"point_to_multipoint": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-multipoint network",
																ValidateFunc:  validation.IntBetween(0, 1),
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
									"address_type": {
										Type: schema.TypeString, Optional: true, Description: "'anycast': Configure an IPv6 anycast address; 'link-local': Configure an IPv6 link local address;",
										ValidateFunc: validation.StringInSlice([]string{"anycast", "link-local"}, false),
									},
								},
							},
						},
						"ipv6_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"v6_acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply ACL rules to incoming packets on this interface (Named Access List)",
							ValidateFunc: validation.StringLenBetween(1, 16),
						},
						"inbound": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ACL applied on incoming packets to this interface",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as NAT inside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as NAT outside",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"ttl_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"router_adver": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Router Advertisements on this interface; 'disable': Disable Router Advertisements on this interface;",
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
									},
									"default_lifetime": {
										Type: schema.TypeInt, Optional: true, Default: 1800, Description: "Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))",
										ValidateFunc: validation.IntBetween(0, 9000),
									},
									"hop_limit": {
										Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set Router Advertisement Hop Limit (default: 255)",
										ValidateFunc: validation.IntBetween(0, 255),
									},
									"max_interval": {
										Type: schema.TypeInt, Optional: true, Default: 600, Description: "Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))",
										ValidateFunc: validation.IntBetween(4, 1800),
									},
									"min_interval": {
										Type: schema.TypeInt, Optional: true, Default: 200, Description: "Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))",
										ValidateFunc: validation.IntBetween(3, 1350),
									},
									"rate_limit": {
										Type: schema.TypeInt, Optional: true, Default: 100000, Description: "Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)",
										ValidateFunc: validation.IntBetween(1, 100000),
									},
									"reachable_time": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))",
										ValidateFunc: validation.IntBetween(0, 3600000),
									},
									"retransmit_timer": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set Router Advertisement Retransmit Timer (default: 0)",
										ValidateFunc: validation.IntBetween(0, 4294967295),
									},
									"adver_mtu_disable": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Disable Router Advertisement MTU Option",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"ipv6.0.router_adver.0.adver_mtu"},
									},
									"adver_mtu": {
										Type: schema.TypeInt, Optional: true, Description: "Set Router Advertisement MTU Option",
										ConflictsWith: []string{"ipv6.0.router_adver.0.adver_mtu_disable"},
									},
									"prefix_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"prefix": {
													Type: schema.TypeString, Optional: true, Description: "Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)",
													ValidateFunc: axapi.IsIpv6AddressPlen,
												},
												"not_autonomous": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify that the Prefix is not usable for autoconfiguration (default:autonomous)",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"not_on_link": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify that the Prefix is not On-Link (default: on-link)",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"preferred_lifetime": {
													Type: schema.TypeInt, Optional: true, Default: 604800, Description: "Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))",
													ValidateFunc: validation.IntBetween(0, 4294967295),
												},
												"valid_lifetime": {
													Type: schema.TypeInt, Optional: true, Default: 2592000, Description: "Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))",
													ValidateFunc: validation.IntBetween(0, 4294967295),
												},
											},
										},
									},
									"managed_config_action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Managed Address Configuration flag; 'disable': Disable the Managed Address Configuration flag (default);",
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
									},
									"other_config_action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable the Other Stateful Configuration flag; 'disable': Disable the Other Stateful Configuration flag (default);",
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
									},
									"adver_vrid": {
										Type: schema.TypeInt, Optional: true, Description: "Vrid",
										ValidateFunc:  validation.IntBetween(1, 31),
										ConflictsWith: []string{"ipv6.0.router_adver.0.adver_vrid_default"},
									},
									"use_floating_ip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a floating IP as the source address for Router advertisements",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"floating_ip": {
										Type: schema.TypeString, Optional: true, Description: "Use a floating IP as the source address for Router advertisements",
										ValidateFunc: validation.IsIPv6Address,
									},
									"adver_vrid_default": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default VRRP-A vrid",
										ValidateFunc:  validation.IntBetween(0, 1),
										ConflictsWith: []string{"ipv6.0.router_adver.0.adver_vrid"},
									},
									"use_floating_ip_default_vrid": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a floating IP as the source address for Router advertisements",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"floating_ip_default_vrid": {
										Type: schema.TypeString, Optional: true, Description: "Use a floating IP as the source address for Router advertisements",
										ValidateFunc: validation.IsIPv6Address,
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"stateful_firewall": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inside (private) interface for stateful firewall",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"class_list": {
										Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"outside": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outside (public) interface for stateful firewall",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"access_list": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Access-list for traffic from the outside",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"acl_name": {
										Type: schema.TypeString, Optional: true, Description: "Access-list Name",
										ValidateFunc: validation.StringLenBetween(1, 16),
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
			"l3_vlan_fwd_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable L3 forwarding between VLANs for incoming packets on this interface",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"lw_4o6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure LW-4over6 inside interface",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure LW-4over6 outside interface",
							ValidateFunc: validation.IntBetween(0, 1),
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
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP outside interface",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"map_t_inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP inside interface (connected to MAP domains)",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"map_t_outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP outside interface",
							ValidateFunc: validation.IntBetween(0, 1),
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
			"nptv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_name": {
										Type: schema.TypeString, Required: true, Description: "NPTv6 domain name",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"bind_type": {
										Type: schema.TypeString, Required: true, Description: "'inside': This interface is connected to NPTv6 inside networks; 'outside': This interface is connected to NPTv6 outside networks;",
										ValidateFunc: validation.StringInSlice([]string{"inside", "outside"}, false),
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
			"ping_sweep_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
			},
			"port_scan_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable port scan detection; 'disable': Disable port scan detection(default);",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num_pkts': Input packets; 'num_total_bytes': Input bytes; 'num_unicast_pkts': Received unicasts; 'num_broadcast_pkts': Received broadcasts; 'num_multicast_pkts': Received multicasts; 'num_tx_pkts': Transmitted packets; 'num_total_tx_bytes': Transmitted bytes; 'num_unicast_tx_pkts': Transmitted unicasts; 'num_broadcast_tx_pkts': Transmitted broadcasts; 'num_multicast_tx_pkts': Transmitted multicasts; 'rate_pkt_sent': Packet sent rate packets/sec; 'rate_byte_sent': Byte sent rate bits/sec; 'rate_pkt_rcvd': Packet received rate packets/sec; 'rate_byte_rcvd': Byte received rate bits/sec; 'load_interval': Load Interval;",
							ValidateFunc: validation.StringInSlice([]string{"all", "num_pkts", "num_total_bytes", "num_unicast_pkts", "num_broadcast_pkts", "num_multicast_pkts", "num_tx_pkts", "num_total_tx_bytes", "num_unicast_tx_pkts", "num_broadcast_tx_pkts", "num_multicast_tx_pkts", "rate_pkt_sent", "rate_byte_sent", "rate_pkt_rcvd", "rate_byte_rcvd", "load_interval"}, false),
						},
					},
				},
			},
			"trap_source": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The trap source",
				ValidateFunc: validation.IntBetween(0, 1),
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

func resourceInterfaceVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVe(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVe(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVe(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVe(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceVeAccessList(d []interface{}) edpt.InterfaceVeAccessList {
	var ret edpt.InterfaceVeAccessList
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

func getObjectInterfaceVeBfd(d []interface{}) edpt.InterfaceVeBfd {
	var ret edpt.InterfaceVeBfd
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceVeBfdAuthentication(in["authentication"].([]interface{}))
		ret.Echo = in["echo"].(int)
		ret.Demand = in["demand"].(int)
		ret.IntervalCfg = getObjectInterfaceVeBfdIntervalCfg(in["interval_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeBfdAuthentication(d []interface{}) edpt.InterfaceVeBfdAuthentication {
	var ret edpt.InterfaceVeBfdAuthentication
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

func getObjectInterfaceVeBfdIntervalCfg(d []interface{}) edpt.InterfaceVeBfdIntervalCfg {
	var ret edpt.InterfaceVeBfdIntervalCfg
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

func getObjectInterfaceVeDdos(d []interface{}) edpt.InterfaceVeDdos {
	var ret edpt.InterfaceVeDdos
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIcmpRateLimit(d []interface{}) edpt.InterfaceVeIcmpRateLimit {
	var ret edpt.InterfaceVeIcmpRateLimit
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Normal = in["normal"].(int)
		ret.Lockup = in["lockup"].(int)
		ret.LockupPeriod = in["lockup_period"].(int)
	}
	return ret
}

func getObjectInterfaceVeIcmpv6RateLimit(d []interface{}) edpt.InterfaceVeIcmpv6RateLimit {
	var ret edpt.InterfaceVeIcmpv6RateLimit
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NormalV6 = in["normal_v6"].(int)
		ret.LockupV6 = in["lockup_v6"].(int)
		ret.LockupPeriodV6 = in["lockup_period_v6"].(int)
	}
	return ret
}

func getObjectInterfaceVeIp(d []interface{}) edpt.InterfaceVeIp {
	var ret edpt.InterfaceVeIp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.AddressList = getSliceInterfaceVeIpAddressList(in["address_list"].([]interface{}))
		ret.AllowPromiscuousVip = in["allow_promiscuous_vip"].(int)
		ret.Client = in["client"].(int)
		ret.Server = in["server"].(int)
		ret.HelperAddressList = getSliceInterfaceVeIpHelperAddressList(in["helper_address_list"].([]interface{}))
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.SynCookie = in["syn_cookie"].(int)
		ret.SlbPartitionRedirect = in["slb_partition_redirect"].(int)
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.QueryInterval = in["query_interval"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceVeIpStatefulFirewall(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceVeIpRouter(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceVeIpRip(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceVeIpOspf(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeIpAddressList(d []interface{}) []edpt.InterfaceVeIpAddressList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpAddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpHelperAddressList(d []interface{}) []edpt.InterfaceVeIpHelperAddressList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpHelperAddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpHelperAddressList
		oi.HelperAddress = in["helper_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpStatefulFirewall(d []interface{}) edpt.InterfaceVeIpStatefulFirewall {
	var ret edpt.InterfaceVeIpStatefulFirewall
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.Outside = in["outside"].(int)
		ret.AccessList = in["access_list"].(int)
		ret.AclId = in["acl_id"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpRouter(d []interface{}) edpt.InterfaceVeIpRouter {
	var ret edpt.InterfaceVeIpRouter
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Isis = getObjectInterfaceVeIpRouterIsis(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpRouterIsis(d []interface{}) edpt.InterfaceVeIpRouterIsis {
	var ret edpt.InterfaceVeIpRouterIsis
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

func getObjectInterfaceVeIpRip(d []interface{}) edpt.InterfaceVeIpRip {
	var ret edpt.InterfaceVeIpRip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceVeIpRipAuthentication(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceVeIpRipSendCfg(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceVeIpRipReceiveCfg(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceVeIpRipSplitHorizonCfg(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthentication(d []interface{}) edpt.InterfaceVeIpRipAuthentication {
	var ret edpt.InterfaceVeIpRipAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Str = getObjectInterfaceVeIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceVeIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceVeIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationStr(d []interface{}) edpt.InterfaceVeIpRipAuthenticationStr {
	var ret edpt.InterfaceVeIpRipAuthenticationStr
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationMode(d []interface{}) edpt.InterfaceVeIpRipAuthenticationMode {
	var ret edpt.InterfaceVeIpRipAuthenticationMode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceVeIpRipAuthenticationKeyChain {
	var ret edpt.InterfaceVeIpRipAuthenticationKeyChain
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipSendCfg(d []interface{}) edpt.InterfaceVeIpRipSendCfg {
	var ret edpt.InterfaceVeIpRipSendCfg
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

func getObjectInterfaceVeIpRipReceiveCfg(d []interface{}) edpt.InterfaceVeIpRipReceiveCfg {
	var ret edpt.InterfaceVeIpRipReceiveCfg
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

func getObjectInterfaceVeIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceVeIpRipSplitHorizonCfg {
	var ret edpt.InterfaceVeIpRipSplitHorizonCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpOspf(d []interface{}) edpt.InterfaceVeIpOspf {
	var ret edpt.InterfaceVeIpOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceVeIpOspfOspfGlobal(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceVeIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobal(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobal {
	var ret edpt.InterfaceVeIpOspfOspfGlobal
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceVeIpOspfOspfGlobalAuthenticationCfg(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceVeIpOspfOspfGlobalBfdCfg(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceVeIpOspfOspfGlobalDatabaseFilterCfg(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceVeIpOspfOspfGlobalMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceVeIpOspfOspfGlobalNetwork(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalAuthenticationCfg(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalAuthenticationCfg {
	var ret edpt.InterfaceVeIpOspfOspfGlobalAuthenticationCfg
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

func getObjectInterfaceVeIpOspfOspfGlobalBfdCfg(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalBfdCfg {
	var ret edpt.InterfaceVeIpOspfOspfGlobalBfdCfg
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

func getObjectInterfaceVeIpOspfOspfGlobalDatabaseFilterCfg(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg {
	var ret edpt.InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg
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

func getSliceInterfaceVeIpOspfOspfGlobalMessageDigestCfg(d []interface{}) []edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5 {
	var ret edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5
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

func getObjectInterfaceVeIpOspfOspfGlobalNetwork(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalNetwork {
	var ret edpt.InterfaceVeIpOspfOspfGlobalNetwork
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

func getSliceInterfaceVeIpOspfOspfIpList(d []interface{}) []edpt.InterfaceVeIpOspfOspfIpList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpOspfOspfIpList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceVeIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceVeIpOspfOspfIpListMessageDigestCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpOspfOspfIpListMessageDigestCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6(d []interface{}) edpt.InterfaceVeIpv6 {
	var ret edpt.InterfaceVeIpv6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AddressList = getSliceInterfaceVeIpv6AddressList(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.V6AclName = in["v6_acl_name"].(string)
		ret.Inbound = in["inbound"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.RouterAdver = getObjectInterfaceVeIpv6RouterAdver(in["router_adver"].([]interface{}))
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceVeIpv6StatefulFirewall(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceVeIpv6Router(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceVeIpv6Rip(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceVeIpv6Ospf(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeIpv6AddressList(d []interface{}) []edpt.InterfaceVeIpv6AddressList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6AddressList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterAdver(d []interface{}) edpt.InterfaceVeIpv6RouterAdver {
	var ret edpt.InterfaceVeIpv6RouterAdver
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.DefaultLifetime = in["default_lifetime"].(int)
		ret.HopLimit = in["hop_limit"].(int)
		ret.MaxInterval = in["max_interval"].(int)
		ret.MinInterval = in["min_interval"].(int)
		ret.RateLimit = in["rate_limit"].(int)
		ret.ReachableTime = in["reachable_time"].(int)
		ret.RetransmitTimer = in["retransmit_timer"].(int)
		ret.AdverMtuDisable = in["adver_mtu_disable"].(int)
		ret.AdverMtu = in["adver_mtu"].(int)
		ret.PrefixList = getSliceInterfaceVeIpv6RouterAdverPrefixList(in["prefix_list"].([]interface{}))
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

func getSliceInterfaceVeIpv6RouterAdverPrefixList(d []interface{}) []edpt.InterfaceVeIpv6RouterAdverPrefixList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6RouterAdverPrefixList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6RouterAdverPrefixList
		oi.Prefix = in["prefix"].(string)
		oi.NotAutonomous = in["not_autonomous"].(int)
		oi.NotOnLink = in["not_on_link"].(int)
		oi.PreferredLifetime = in["preferred_lifetime"].(int)
		oi.ValidLifetime = in["valid_lifetime"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6StatefulFirewall(d []interface{}) edpt.InterfaceVeIpv6StatefulFirewall {
	var ret edpt.InterfaceVeIpv6StatefulFirewall
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.Outside = in["outside"].(int)
		ret.AccessList = in["access_list"].(int)
		ret.AclName = in["acl_name"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpv6Router(d []interface{}) edpt.InterfaceVeIpv6Router {
	var ret edpt.InterfaceVeIpv6Router
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ripng = getObjectInterfaceVeIpv6RouterRipng(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceVeIpv6RouterOspf(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceVeIpv6RouterIsis(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterRipng(d []interface{}) edpt.InterfaceVeIpv6RouterRipng {
	var ret edpt.InterfaceVeIpv6RouterRipng
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

func getObjectInterfaceVeIpv6RouterOspf(d []interface{}) edpt.InterfaceVeIpv6RouterOspf {
	var ret edpt.InterfaceVeIpv6RouterOspf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.AreaList = getSliceInterfaceVeIpv6RouterOspfAreaList(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceVeIpv6RouterOspfAreaList(d []interface{}) []edpt.InterfaceVeIpv6RouterOspfAreaList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6RouterOspfAreaList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6RouterOspfAreaList
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterIsis(d []interface{}) edpt.InterfaceVeIpv6RouterIsis {
	var ret edpt.InterfaceVeIpv6RouterIsis
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

func getObjectInterfaceVeIpv6Rip(d []interface{}) edpt.InterfaceVeIpv6Rip {
	var ret edpt.InterfaceVeIpv6Rip
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceVeIpv6RipSplitHorizonCfg(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpv6RipSplitHorizonCfg(d []interface{}) edpt.InterfaceVeIpv6RipSplitHorizonCfg {
	var ret edpt.InterfaceVeIpv6RipSplitHorizonCfg
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpv6Ospf(d []interface{}) edpt.InterfaceVeIpv6Ospf {
	var ret edpt.InterfaceVeIpv6Ospf
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.NetworkList = getSliceInterfaceVeIpv6OspfNetworkList(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceVeIpv6OspfCostCfg(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceVeIpv6OspfDeadIntervalCfg(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceVeIpv6OspfHelloIntervalCfg(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceVeIpv6OspfMtuIgnoreCfg(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceVeIpv6OspfNeighborCfg(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceVeIpv6OspfPriorityCfg(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceVeIpv6OspfRetransmitIntervalCfg(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceVeIpv6OspfTransmitDelayCfg(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfNetworkList(d []interface{}) []edpt.InterfaceVeIpv6OspfNetworkList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfNetworkList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfNetworkList
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfCostCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfCostCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfDeadIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfDeadIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfHelloIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfHelloIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfMtuIgnoreCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfMtuIgnoreCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfNeighborCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfNeighborCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfNeighborCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfNeighborCfg
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfPriorityCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfPriorityCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfTransmitDelayCfg {
	count := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfTransmitDelayCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIsis(d []interface{}) edpt.InterfaceVeIsis {
	var ret edpt.InterfaceVeIsis
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Authentication = getObjectInterfaceVeIsisAuthentication(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceVeIsisBfdCfg(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceVeIsisCsnpIntervalList(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceVeIsisHelloIntervalList(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceVeIsisHelloIntervalMinimalList(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceVeIsisHelloMultiplierList(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceVeIsisMeshGroup(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceVeIsisMetricList(in["metric_list"].([]interface{}))
		ret.Network = in["network"].(string)
		ret.PasswordList = getSliceInterfaceVeIsisPasswordList(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceVeIsisPriorityList(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceVeIsisWideMetricList(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIsisAuthentication(d []interface{}) edpt.InterfaceVeIsisAuthentication {
	var ret edpt.InterfaceVeIsisAuthentication
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceVeIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceVeIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceVeIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceVeIsisAuthenticationSendOnlyList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisAuthenticationSendOnlyList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceVeIsisAuthenticationModeList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisAuthenticationModeList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceVeIsisAuthenticationKeyChainList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisAuthenticationKeyChainList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIsisBfdCfg(d []interface{}) edpt.InterfaceVeIsisBfdCfg {
	var ret edpt.InterfaceVeIsisBfdCfg
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

func getSliceInterfaceVeIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceVeIsisCsnpIntervalList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisCsnpIntervalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisHelloIntervalList(d []interface{}) []edpt.InterfaceVeIsisHelloIntervalList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisHelloIntervalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceVeIsisHelloIntervalMinimalList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisHelloIntervalMinimalList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceVeIsisHelloMultiplierList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisHelloMultiplierList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIsisMeshGroup(d []interface{}) edpt.InterfaceVeIsisMeshGroup {
	var ret edpt.InterfaceVeIsisMeshGroup
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

func getSliceInterfaceVeIsisMetricList(d []interface{}) []edpt.InterfaceVeIsisMetricList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisMetricList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisPasswordList(d []interface{}) []edpt.InterfaceVeIsisPasswordList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisPasswordList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisPriorityList(d []interface{}) []edpt.InterfaceVeIsisPriorityList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisPriorityList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisWideMetricList(d []interface{}) []edpt.InterfaceVeIsisWideMetricList {
	count := len(d)
	ret := make([]edpt.InterfaceVeIsisWideMetricList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeLw4o6(d []interface{}) edpt.InterfaceVeLw4o6 {
	var ret edpt.InterfaceVeLw4o6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeMap(d []interface{}) edpt.InterfaceVeMap {
	var ret edpt.InterfaceVeMap
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.MapTInside = in["map_t_inside"].(int)
		ret.MapTOutside = in["map_t_outside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeNptv6(d []interface{}) edpt.InterfaceVeNptv6 {
	var ret edpt.InterfaceVeNptv6
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.DomainList = getSliceInterfaceVeNptv6DomainList(in["domain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeNptv6DomainList(d []interface{}) []edpt.InterfaceVeNptv6DomainList {
	count := len(d)
	ret := make([]edpt.InterfaceVeNptv6DomainList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeNptv6DomainList
		oi.DomainName = in["domain_name"].(string)
		oi.BindType = in["bind_type"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeSamplingEnable(d []interface{}) []edpt.InterfaceVeSamplingEnable {
	count := len(d)
	ret := make([]edpt.InterfaceVeSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceVe(d *schema.ResourceData) edpt.InterfaceVe {
	var ret edpt.InterfaceVe
	ret.Inst.AccessList = getObjectInterfaceVeAccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Bfd = getObjectInterfaceVeBfd(d.Get("bfd").([]interface{}))
	ret.Inst.Ddos = getObjectInterfaceVeDdos(d.Get("ddos").([]interface{}))
	ret.Inst.IcmpRateLimit = getObjectInterfaceVeIcmpRateLimit(d.Get("icmp_rate_limit").([]interface{}))
	ret.Inst.Icmpv6RateLimit = getObjectInterfaceVeIcmpv6RateLimit(d.Get("icmpv6_rate_limit").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceVeIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceVeIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceVeIsis(d.Get("isis").([]interface{}))
	ret.Inst.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	ret.Inst.Lw4o6 = getObjectInterfaceVeLw4o6(d.Get("lw_4o6").([]interface{}))
	ret.Inst.Map = getObjectInterfaceVeMap(d.Get("map").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Nptv6 = getObjectInterfaceVeNptv6(d.Get("nptv6").([]interface{}))
	ret.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	ret.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	ret.Inst.SamplingEnable = getSliceInterfaceVeSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.TrapSource = d.Get("trap_source").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
