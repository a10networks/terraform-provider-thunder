package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunk() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk`: Trunk interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkCreate,
		UpdateContext: resourceInterfaceTrunkUpdate,
		ReadContext:   resourceInterfaceTrunkRead,
		DeleteContext: resourceInterfaceTrunkDelete,

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
						"per_member_port": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_address": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 local-address",
									},
									"neighbor_address": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 neighbor address",
									},
									"ipv6_local": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 local-address",
									},
									"ipv6_nbr": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 neighbor-address",
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
			"ddos": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DDoS inside (trusted) or outside (untrusted) interface",
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DDoS inside (trusted) or outside (untrusted) interface",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"do_auto_recovery": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "(Only for LACP trunks) Attempt auto-recovery after ports-treshold is triggered",
			},
			"gaming_protocol_compliance": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Gaming Protocol Compliance Check",
			},
			"icmp_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"normal": {
							Type: schema.TypeInt, Optional: true, Description: "Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit",
						},
						"lockup": {
							Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
						},
						"lockup_period": {
							Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
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
						},
						"lockup_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
						},
						"lockup_period_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
						},
					},
				},
			},
			"ifnum": {
				Type: schema.TypeInt, Required: true, Description: "Trunk interface number",
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
						"client": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client facing interface for IPv4/v6 traffic",
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server facing interface for IPv4/v6 traffic",
						},
						"cache_spoofing_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "This interface connects to spoofing cache",
						},
						"helper_address_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"helper_address": {
										Type: schema.TypeString, Optional: true, Description: "Helper address for DHCP packets (IP address)",
									},
								},
							},
						},
						"nat": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
									},
									"outside": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
									},
								},
							},
						},
						"ttl_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet",
						},
						"syn_cookie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SYN-cookie on the interface",
						},
						"slb_partition_redirect": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Redirect SLB traffic across partition",
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
									"acl_id": {
										Type: schema.TypeInt, Optional: true, Description: "ACL id",
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
									"address_type": {
										Type: schema.TypeString, Optional: true, Description: "'anycast': Configure an IPv6 anycast address; 'link-local': Configure an IPv6 link local address;",
									},
								},
							},
						},
						"ipv6_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
						},
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
						"ttl_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"l3_vlan_fwd_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable L3 forwarding between VLANs",
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
			"mac_learning": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable MAC learning; 'disable': Disable MAC learning; 'dmac-only': Enable destination MAC learning only;",
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
									},
									"bind_type": {
										Type: schema.TypeString, Required: true, Description: "'inside': This interface is connected to NPTv6 inside networks; 'outside': This interface is connected to NPTv6 outside networks;",
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
			"ports_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold for the minimum number of ports that need to be UP for the trunk to remain UP",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num_pkts': num_pkts; 'num_total_bytes': num_total_bytes; 'num_unicast_pkts': num_unicast_pkts; 'num_broadcast_pkts': num_broadcast_pkts; 'num_multicast_pkts': num_multicast_pkts; 'num_tx_pkts': num_tx_pkts; 'num_total_tx_bytes': num_total_tx_bytes; 'num_unicast_tx_pkts': num_unicast_tx_pkts; 'num_broadcast_tx_pkts': num_broadcast_tx_pkts; 'num_multicast_tx_pkts': num_multicast_tx_pkts; 'dropped_dis_rx_pkts': dropped_dis_rx_pkts; 'dropped_rx_pkts': dropped_rx_pkts; 'dropped_dis_tx_pkts': dropped_dis_tx_pkts; 'dropped_tx_pkts': dropped_tx_pkts;",
						},
					},
				},
			},
			"spanning_tree": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_edge": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable auto-edge",
						},
						"admin_edge": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable admin-edge",
						},
						"instance_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_start": {
										Type: schema.TypeInt, Optional: true, Description: "Instance ID",
									},
									"mstp_path_cost": {
										Type: schema.TypeInt, Optional: true, Description: "Path cost (Limit)",
									},
								},
							},
						},
						"path_cost": {
							Type: schema.TypeInt, Optional: true, Description: "Path cost (Limit)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sync_modify_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable SYNC bit modify for ports-threshold do-auto-recovery",
			},
			"timer": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Timer to re-check the threshold under certain conditions (Time in seconds (Default: 10))",
			},
			"trap_source": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The trap source",
			},
			"update_l2_info": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Update and use received L2 information",
			},
			"use_hw_hash": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable HW based load balacing decision rule",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_wire": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark trunk as a virtual wire interface",
			},
			"vlan_learning": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable VLAN learning; 'disable': Disable VLAN learning;",
			},
		},
	}
}
func resourceInterfaceTrunkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunk(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunk(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunk(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunk(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTrunkAccessList(d []interface{}) edpt.InterfaceTrunkAccessList {

	count1 := len(d)
	var ret edpt.InterfaceTrunkAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkBfd780(d []interface{}) edpt.InterfaceTrunkBfd780 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfd780
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceTrunkBfdAuthentication781(in["authentication"].([]interface{}))
		ret.Echo = in["echo"].(int)
		ret.Demand = in["demand"].(int)
		ret.IntervalCfg = getObjectInterfaceTrunkBfdIntervalCfg782(in["interval_cfg"].([]interface{}))
		//omit uuid
		ret.PerMemberPort = getObjectInterfaceTrunkBfdPerMemberPort783(in["per_member_port"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkBfdAuthentication781(d []interface{}) edpt.InterfaceTrunkBfdAuthentication781 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfdAuthentication781
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceTrunkBfdIntervalCfg782(d []interface{}) edpt.InterfaceTrunkBfdIntervalCfg782 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfdIntervalCfg782
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkBfdPerMemberPort783(d []interface{}) edpt.InterfaceTrunkBfdPerMemberPort783 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkBfdPerMemberPort783
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LocalAddress = in["local_address"].(string)
		ret.NeighborAddress = in["neighbor_address"].(string)
		ret.Ipv6Local = in["ipv6_local"].(string)
		ret.Ipv6Nbr = in["ipv6_nbr"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkDdos784(d []interface{}) edpt.InterfaceTrunkDdos784 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkDdos784
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIcmpRateLimit(d []interface{}) edpt.InterfaceTrunkIcmpRateLimit {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIcmpRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Normal = in["normal"].(int)
		ret.Lockup = in["lockup"].(int)
		ret.LockupPeriod = in["lockup_period"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIcmpv6RateLimit(d []interface{}) edpt.InterfaceTrunkIcmpv6RateLimit {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIcmpv6RateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NormalV6 = in["normal_v6"].(int)
		ret.LockupV6 = in["lockup_v6"].(int)
		ret.LockupPeriodV6 = in["lockup_period_v6"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIp785(d []interface{}) edpt.InterfaceTrunkIp785 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIp785
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.AddressList = getSliceInterfaceTrunkIpAddressList786(in["address_list"].([]interface{}))
		ret.AllowPromiscuousVip = in["allow_promiscuous_vip"].(int)
		ret.Client = in["client"].(int)
		ret.Server = in["server"].(int)
		ret.CacheSpoofingPort = in["cache_spoofing_port"].(int)
		ret.HelperAddressList = getSliceInterfaceTrunkIpHelperAddressList787(in["helper_address_list"].([]interface{}))
		ret.Nat = getObjectInterfaceTrunkIpNat788(in["nat"].([]interface{}))
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.SynCookie = in["syn_cookie"].(int)
		ret.SlbPartitionRedirect = in["slb_partition_redirect"].(int)
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.QueryInterval = in["query_interval"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		ret.Unnumbered = in["unnumbered"].(int)
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceTrunkIpStatefulFirewall789(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceTrunkIpRouter790(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceTrunkIpRip792(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceTrunkIpOspf800(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTrunkIpAddressList786(d []interface{}) []edpt.InterfaceTrunkIpAddressList786 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpAddressList786, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpAddressList786
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpHelperAddressList787(d []interface{}) []edpt.InterfaceTrunkIpHelperAddressList787 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpHelperAddressList787, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpHelperAddressList787
		oi.HelperAddress = in["helper_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpNat788(d []interface{}) edpt.InterfaceTrunkIpNat788 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpNat788
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpStatefulFirewall789(d []interface{}) edpt.InterfaceTrunkIpStatefulFirewall789 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpStatefulFirewall789
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.Outside = in["outside"].(int)
		ret.AccessList = in["access_list"].(int)
		ret.AclId = in["acl_id"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpRouter790(d []interface{}) edpt.InterfaceTrunkIpRouter790 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRouter790
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceTrunkIpRouterIsis791(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpRouterIsis791(d []interface{}) edpt.InterfaceTrunkIpRouterIsis791 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRouterIsis791
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpRip792(d []interface{}) edpt.InterfaceTrunkIpRip792 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRip792
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceTrunkIpRipAuthentication793(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceTrunkIpRipSendCfg797(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceTrunkIpRipReceiveCfg798(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceTrunkIpRipSplitHorizonCfg799(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthentication793(d []interface{}) edpt.InterfaceTrunkIpRipAuthentication793 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthentication793
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceTrunkIpRipAuthenticationStr794(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceTrunkIpRipAuthenticationMode795(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceTrunkIpRipAuthenticationKeyChain796(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationStr794(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationStr794 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationStr794
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationMode795(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationMode795 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationMode795
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationKeyChain796(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationKeyChain796 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationKeyChain796
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipSendCfg797(d []interface{}) edpt.InterfaceTrunkIpRipSendCfg797 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipSendCfg797
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipReceiveCfg798(d []interface{}) edpt.InterfaceTrunkIpRipReceiveCfg798 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipReceiveCfg798
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipSplitHorizonCfg799(d []interface{}) edpt.InterfaceTrunkIpRipSplitHorizonCfg799 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipSplitHorizonCfg799
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspf800(d []interface{}) edpt.InterfaceTrunkIpOspf800 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspf800
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceTrunkIpOspfOspfGlobal801(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceTrunkIpOspfOspfIpList808(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobal801(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobal801 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobal801
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceTrunkIpOspfOspfGlobalAuthenticationCfg802(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceTrunkIpOspfOspfGlobalBfdCfg803(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg804(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceTrunkIpOspfOspfGlobalNetwork807(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalAuthenticationCfg802(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg802 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg802
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalBfdCfg803(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalBfdCfg803 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalBfdCfg803
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg804(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg804 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg804
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805(d []interface{}) []edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5806(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5806(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5806 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5806
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalNetwork807(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalNetwork807 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalNetwork807
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

func getSliceInterfaceTrunkIpOspfOspfIpList808(d []interface{}) []edpt.InterfaceTrunkIpOspfOspfIpList808 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpOspfOspfIpList808, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpOspfOspfIpList808
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceTrunkIpOspfOspfIpListMessageDigestCfg809(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpOspfOspfIpListMessageDigestCfg809(d []interface{}) []edpt.InterfaceTrunkIpOspfOspfIpListMessageDigestCfg809 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpOspfOspfIpListMessageDigestCfg809, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpOspfOspfIpListMessageDigestCfg809
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6810(d []interface{}) edpt.InterfaceTrunkIpv6810 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6810
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceInterfaceTrunkIpv6AddressList811(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.AccessListCfg = getObjectInterfaceTrunkIpv6AccessListCfg812(in["access_list_cfg"].([]interface{}))
		ret.Nat = getObjectInterfaceTrunkIpv6Nat813(in["nat"].([]interface{}))
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.RouterAdver = getObjectInterfaceTrunkIpv6RouterAdver814(in["router_adver"].([]interface{}))
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceTrunkIpv6StatefulFirewall818(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceTrunkIpv6Router819(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceTrunkIpv6Rip824(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceTrunkIpv6Ospf826(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTrunkIpv6AddressList811(d []interface{}) []edpt.InterfaceTrunkIpv6AddressList811 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6AddressList811, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6AddressList811
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6AccessListCfg812(d []interface{}) edpt.InterfaceTrunkIpv6AccessListCfg812 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6AccessListCfg812
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V6AclName = in["v6_acl_name"].(string)
		ret.Inbound = in["inbound"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Nat813(d []interface{}) edpt.InterfaceTrunkIpv6Nat813 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Nat813
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterAdver814(d []interface{}) edpt.InterfaceTrunkIpv6RouterAdver814 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterAdver814
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
		ret.Mtu = getObjectInterfaceTrunkIpv6RouterAdverMtu815(in["mtu"].([]interface{}))
		ret.PrefixList = getSliceInterfaceTrunkIpv6RouterAdverPrefixList816(in["prefix_list"].([]interface{}))
		ret.ManagedConfigAction = in["managed_config_action"].(string)
		ret.OtherConfigAction = in["other_config_action"].(string)
		ret.Vrid = getObjectInterfaceTrunkIpv6RouterAdverVrid817(in["vrid"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterAdverMtu815(d []interface{}) edpt.InterfaceTrunkIpv6RouterAdverMtu815 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterAdverMtu815
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdverMtuDisable = in["adver_mtu_disable"].(int)
		ret.AdverMtu = in["adver_mtu"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6RouterAdverPrefixList816(d []interface{}) []edpt.InterfaceTrunkIpv6RouterAdverPrefixList816 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6RouterAdverPrefixList816, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6RouterAdverPrefixList816
		oi.Prefix = in["prefix"].(string)
		oi.NotAutonomous = in["not_autonomous"].(int)
		oi.NotOnLink = in["not_on_link"].(int)
		oi.PreferredLifetime = in["preferred_lifetime"].(int)
		oi.ValidLifetime = in["valid_lifetime"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterAdverVrid817(d []interface{}) edpt.InterfaceTrunkIpv6RouterAdverVrid817 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterAdverVrid817
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

func getObjectInterfaceTrunkIpv6StatefulFirewall818(d []interface{}) edpt.InterfaceTrunkIpv6StatefulFirewall818 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6StatefulFirewall818
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

func getObjectInterfaceTrunkIpv6Router819(d []interface{}) edpt.InterfaceTrunkIpv6Router819 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Router819
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceTrunkIpv6RouterRipng820(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceTrunkIpv6RouterOspf821(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceTrunkIpv6RouterIsis823(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterRipng820(d []interface{}) edpt.InterfaceTrunkIpv6RouterRipng820 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterRipng820
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterOspf821(d []interface{}) edpt.InterfaceTrunkIpv6RouterOspf821 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterOspf821
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceTrunkIpv6RouterOspfAreaList822(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTrunkIpv6RouterOspfAreaList822(d []interface{}) []edpt.InterfaceTrunkIpv6RouterOspfAreaList822 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6RouterOspfAreaList822, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6RouterOspfAreaList822
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RouterIsis823(d []interface{}) edpt.InterfaceTrunkIpv6RouterIsis823 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RouterIsis823
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Rip824(d []interface{}) edpt.InterfaceTrunkIpv6Rip824 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Rip824
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceTrunkIpv6RipSplitHorizonCfg825(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpv6RipSplitHorizonCfg825(d []interface{}) edpt.InterfaceTrunkIpv6RipSplitHorizonCfg825 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RipSplitHorizonCfg825
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpv6Ospf826(d []interface{}) edpt.InterfaceTrunkIpv6Ospf826 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6Ospf826
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceTrunkIpv6OspfNetworkList827(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceTrunkIpv6OspfCostCfg828(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceTrunkIpv6OspfDeadIntervalCfg829(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceTrunkIpv6OspfHelloIntervalCfg830(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceTrunkIpv6OspfMtuIgnoreCfg831(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceTrunkIpv6OspfNeighborCfg832(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceTrunkIpv6OspfPriorityCfg833(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceTrunkIpv6OspfRetransmitIntervalCfg834(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceTrunkIpv6OspfTransmitDelayCfg835(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfNetworkList827(d []interface{}) []edpt.InterfaceTrunkIpv6OspfNetworkList827 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfNetworkList827, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfNetworkList827
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfCostCfg828(d []interface{}) []edpt.InterfaceTrunkIpv6OspfCostCfg828 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfCostCfg828, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfCostCfg828
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfDeadIntervalCfg829(d []interface{}) []edpt.InterfaceTrunkIpv6OspfDeadIntervalCfg829 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfDeadIntervalCfg829, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfDeadIntervalCfg829
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfHelloIntervalCfg830(d []interface{}) []edpt.InterfaceTrunkIpv6OspfHelloIntervalCfg830 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfHelloIntervalCfg830, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfHelloIntervalCfg830
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfMtuIgnoreCfg831(d []interface{}) []edpt.InterfaceTrunkIpv6OspfMtuIgnoreCfg831 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfMtuIgnoreCfg831, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfMtuIgnoreCfg831
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfNeighborCfg832(d []interface{}) []edpt.InterfaceTrunkIpv6OspfNeighborCfg832 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfNeighborCfg832, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfNeighborCfg832
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfPriorityCfg833(d []interface{}) []edpt.InterfaceTrunkIpv6OspfPriorityCfg833 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfPriorityCfg833, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfPriorityCfg833
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfRetransmitIntervalCfg834(d []interface{}) []edpt.InterfaceTrunkIpv6OspfRetransmitIntervalCfg834 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfRetransmitIntervalCfg834, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfRetransmitIntervalCfg834
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpv6OspfTransmitDelayCfg835(d []interface{}) []edpt.InterfaceTrunkIpv6OspfTransmitDelayCfg835 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6OspfTransmitDelayCfg835, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6OspfTransmitDelayCfg835
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIsis836(d []interface{}) edpt.InterfaceTrunkIsis836 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsis836
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceTrunkIsisAuthentication837(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceTrunkIsisBfdCfg841(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceTrunkIsisCsnpIntervalList842(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceTrunkIsisHelloIntervalList843(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceTrunkIsisHelloIntervalMinimalList844(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceTrunkIsisHelloMultiplierList845(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceTrunkIsisMeshGroup846(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceTrunkIsisMetricList847(in["metric_list"].([]interface{}))
		ret.Network = in["network"].(string)
		ret.PasswordList = getSliceInterfaceTrunkIsisPasswordList848(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceTrunkIsisPriorityList849(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceTrunkIsisWideMetricList850(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIsisAuthentication837(d []interface{}) edpt.InterfaceTrunkIsisAuthentication837 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsisAuthentication837
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceTrunkIsisAuthenticationSendOnlyList838(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceTrunkIsisAuthenticationModeList839(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceTrunkIsisAuthenticationKeyChainList840(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTrunkIsisAuthenticationSendOnlyList838(d []interface{}) []edpt.InterfaceTrunkIsisAuthenticationSendOnlyList838 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisAuthenticationSendOnlyList838, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisAuthenticationSendOnlyList838
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisAuthenticationModeList839(d []interface{}) []edpt.InterfaceTrunkIsisAuthenticationModeList839 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisAuthenticationModeList839, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisAuthenticationModeList839
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisAuthenticationKeyChainList840(d []interface{}) []edpt.InterfaceTrunkIsisAuthenticationKeyChainList840 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisAuthenticationKeyChainList840, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisAuthenticationKeyChainList840
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIsisBfdCfg841(d []interface{}) edpt.InterfaceTrunkIsisBfdCfg841 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsisBfdCfg841
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIsisCsnpIntervalList842(d []interface{}) []edpt.InterfaceTrunkIsisCsnpIntervalList842 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisCsnpIntervalList842, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisCsnpIntervalList842
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisHelloIntervalList843(d []interface{}) []edpt.InterfaceTrunkIsisHelloIntervalList843 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisHelloIntervalList843, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisHelloIntervalList843
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisHelloIntervalMinimalList844(d []interface{}) []edpt.InterfaceTrunkIsisHelloIntervalMinimalList844 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisHelloIntervalMinimalList844, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisHelloIntervalMinimalList844
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisHelloMultiplierList845(d []interface{}) []edpt.InterfaceTrunkIsisHelloMultiplierList845 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisHelloMultiplierList845, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisHelloMultiplierList845
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIsisMeshGroup846(d []interface{}) edpt.InterfaceTrunkIsisMeshGroup846 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsisMeshGroup846
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIsisMetricList847(d []interface{}) []edpt.InterfaceTrunkIsisMetricList847 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisMetricList847, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisMetricList847
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisPasswordList848(d []interface{}) []edpt.InterfaceTrunkIsisPasswordList848 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisPasswordList848, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisPasswordList848
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisPriorityList849(d []interface{}) []edpt.InterfaceTrunkIsisPriorityList849 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisPriorityList849, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisPriorityList849
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisWideMetricList850(d []interface{}) []edpt.InterfaceTrunkIsisWideMetricList850 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisWideMetricList850, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisWideMetricList850
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkLw4o6851(d []interface{}) edpt.InterfaceTrunkLw4o6851 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkLw4o6851
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkMap852(d []interface{}) edpt.InterfaceTrunkMap852 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkMap852
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

func getObjectInterfaceTrunkNptv6853(d []interface{}) edpt.InterfaceTrunkNptv6853 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkNptv6853
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainList = getSliceInterfaceTrunkNptv6DomainList(in["domain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTrunkNptv6DomainList(d []interface{}) []edpt.InterfaceTrunkNptv6DomainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkNptv6DomainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkNptv6DomainList
		oi.DomainName = in["domain_name"].(string)
		oi.BindType = in["bind_type"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkSamplingEnable(d []interface{}) []edpt.InterfaceTrunkSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkSpanningTree854(d []interface{}) edpt.InterfaceTrunkSpanningTree854 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkSpanningTree854
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AutoEdge = in["auto_edge"].(int)
		ret.AdminEdge = in["admin_edge"].(int)
		ret.InstanceList = getSliceInterfaceTrunkSpanningTreeInstanceList855(in["instance_list"].([]interface{}))
		ret.PathCost = in["path_cost"].(int)
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTrunkSpanningTreeInstanceList855(d []interface{}) []edpt.InterfaceTrunkSpanningTreeInstanceList855 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkSpanningTreeInstanceList855, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkSpanningTreeInstanceList855
		oi.InstanceStart = in["instance_start"].(int)
		oi.MstpPathCost = in["mstp_path_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTrunk(d *schema.ResourceData) edpt.InterfaceTrunk {
	var ret edpt.InterfaceTrunk
	ret.Inst.AccessList = getObjectInterfaceTrunkAccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Bfd = getObjectInterfaceTrunkBfd780(d.Get("bfd").([]interface{}))
	ret.Inst.Ddos = getObjectInterfaceTrunkDdos784(d.Get("ddos").([]interface{}))
	ret.Inst.DoAutoRecovery = d.Get("do_auto_recovery").(int)
	ret.Inst.GamingProtocolCompliance = d.Get("gaming_protocol_compliance").(int)
	ret.Inst.IcmpRateLimit = getObjectInterfaceTrunkIcmpRateLimit(d.Get("icmp_rate_limit").([]interface{}))
	ret.Inst.Icmpv6RateLimit = getObjectInterfaceTrunkIcmpv6RateLimit(d.Get("icmpv6_rate_limit").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceTrunkIp785(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceTrunkIpv6810(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceTrunkIsis836(d.Get("isis").([]interface{}))
	ret.Inst.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	ret.Inst.Lw4o6 = getObjectInterfaceTrunkLw4o6851(d.Get("lw_4o6").([]interface{}))
	ret.Inst.MacLearning = d.Get("mac_learning").(string)
	ret.Inst.Map = getObjectInterfaceTrunkMap852(d.Get("map").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Nptv6 = getObjectInterfaceTrunkNptv6853(d.Get("nptv6").([]interface{}))
	ret.Inst.PortsThreshold = d.Get("ports_threshold").(int)
	ret.Inst.SamplingEnable = getSliceInterfaceTrunkSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SpanningTree = getObjectInterfaceTrunkSpanningTree854(d.Get("spanning_tree").([]interface{}))
	ret.Inst.SyncModifyDisable = d.Get("sync_modify_disable").(int)
	ret.Inst.Timer = d.Get("timer").(int)
	ret.Inst.TrapSource = d.Get("trap_source").(int)
	ret.Inst.UpdateL2Info = d.Get("update_l2_info").(int)
	ret.Inst.UseHwHash = d.Get("use_hw_hash").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VirtualWire = d.Get("virtual_wire").(int)
	ret.Inst.VlanLearning = d.Get("vlan_learning").(string)
	return ret
}
