package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVe() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve`: Virtual ethernet interface\n\n__PLACEHOLDER__",
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
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Named Access List",
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
				Type: schema.TypeInt, Required: true, Description: "Virtual ethernet interface number",
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
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
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
						"v6_acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply ACL rules to incoming packets on this interface (Named Access List)",
						},
						"inbound": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ACL applied on incoming packets to this interface",
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as NAT inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as NAT outside",
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
										Type: schema.TypeInt, Optional: true, Default: 255, Description: "Set Router Advertisement Hop Limit (default: 255)",
									},
									"max_interval": {
										Type: schema.TypeInt, Optional: true, Default: 600, Description: "Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))",
									},
									"min_interval": {
										Type: schema.TypeInt, Optional: true, Default: 200, Description: "Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))",
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
										Type: schema.TypeInt, Optional: true, Description: "Vrid",
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable L3 forwarding between VLANs for incoming packets on this interface",
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
			"ping_sweep_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);",
			},
			"port_scan_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable port scan detection; 'disable': Disable port scan detection(default);",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'num_pkts': Input packets; 'num_total_bytes': Input bytes; 'num_unicast_pkts': Received unicasts; 'num_broadcast_pkts': Received broadcasts; 'num_multicast_pkts': Received multicasts; 'num_tx_pkts': Transmitted packets; 'num_total_tx_bytes': Transmitted bytes; 'num_unicast_tx_pkts': Transmitted unicasts; 'num_broadcast_tx_pkts': Transmitted broadcasts; 'num_multicast_tx_pkts': Transmitted multicasts; 'rate_pkt_sent': Packet sent rate packets/sec; 'rate_byte_sent': Byte sent rate bits/sec; 'rate_pkt_rcvd': Packet received rate packets/sec; 'rate_byte_rcvd': Byte received rate bits/sec; 'load_interval': Load Interval;",
						},
					},
				},
			},
			"trap_source": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The trap source",
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

func getObjectInterfaceVeAccessList(d []interface{}) edpt.InterfaceVeAccessList {

	count1 := len(d)
	var ret edpt.InterfaceVeAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceVeBfd962(d []interface{}) edpt.InterfaceVeBfd962 {

	count1 := len(d)
	var ret edpt.InterfaceVeBfd962
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceVeBfdAuthentication963(in["authentication"].([]interface{}))
		ret.Echo = in["echo"].(int)
		ret.Demand = in["demand"].(int)
		ret.IntervalCfg = getObjectInterfaceVeBfdIntervalCfg964(in["interval_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeBfdAuthentication963(d []interface{}) edpt.InterfaceVeBfdAuthentication963 {

	count1 := len(d)
	var ret edpt.InterfaceVeBfdAuthentication963
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceVeBfdIntervalCfg964(d []interface{}) edpt.InterfaceVeBfdIntervalCfg964 {

	count1 := len(d)
	var ret edpt.InterfaceVeBfdIntervalCfg964
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func getObjectInterfaceVeDdos965(d []interface{}) edpt.InterfaceVeDdos965 {

	count1 := len(d)
	var ret edpt.InterfaceVeDdos965
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIcmpRateLimit(d []interface{}) edpt.InterfaceVeIcmpRateLimit {

	count1 := len(d)
	var ret edpt.InterfaceVeIcmpRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Normal = in["normal"].(int)
		ret.Lockup = in["lockup"].(int)
		ret.LockupPeriod = in["lockup_period"].(int)
	}
	return ret
}

func getObjectInterfaceVeIcmpv6RateLimit(d []interface{}) edpt.InterfaceVeIcmpv6RateLimit {

	count1 := len(d)
	var ret edpt.InterfaceVeIcmpv6RateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NormalV6 = in["normal_v6"].(int)
		ret.LockupV6 = in["lockup_v6"].(int)
		ret.LockupPeriodV6 = in["lockup_period_v6"].(int)
	}
	return ret
}

func getObjectInterfaceVeIp966(d []interface{}) edpt.InterfaceVeIp966 {

	count1 := len(d)
	var ret edpt.InterfaceVeIp966
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.AddressList = getSliceInterfaceVeIpAddressList967(in["address_list"].([]interface{}))
		ret.AllowPromiscuousVip = in["allow_promiscuous_vip"].(int)
		ret.Client = in["client"].(int)
		ret.Server = in["server"].(int)
		ret.HelperAddressList = getSliceInterfaceVeIpHelperAddressList968(in["helper_address_list"].([]interface{}))
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.SynCookie = in["syn_cookie"].(int)
		ret.SlbPartitionRedirect = in["slb_partition_redirect"].(int)
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.QueryInterval = in["query_interval"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		ret.Unnumbered = in["unnumbered"].(int)
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceVeIpStatefulFirewall969(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceVeIpRouter970(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceVeIpRip972(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceVeIpOspf980(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeIpAddressList967(d []interface{}) []edpt.InterfaceVeIpAddressList967 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpAddressList967, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpAddressList967
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpHelperAddressList968(d []interface{}) []edpt.InterfaceVeIpHelperAddressList968 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpHelperAddressList968, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpHelperAddressList968
		oi.HelperAddress = in["helper_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpStatefulFirewall969(d []interface{}) edpt.InterfaceVeIpStatefulFirewall969 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpStatefulFirewall969
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

func getObjectInterfaceVeIpRouter970(d []interface{}) edpt.InterfaceVeIpRouter970 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRouter970
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceVeIpRouterIsis971(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpRouterIsis971(d []interface{}) edpt.InterfaceVeIpRouterIsis971 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRouterIsis971
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpRip972(d []interface{}) edpt.InterfaceVeIpRip972 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRip972
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceVeIpRipAuthentication973(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceVeIpRipSendCfg977(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceVeIpRipReceiveCfg978(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceVeIpRipSplitHorizonCfg979(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthentication973(d []interface{}) edpt.InterfaceVeIpRipAuthentication973 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthentication973
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceVeIpRipAuthenticationStr974(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceVeIpRipAuthenticationMode975(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceVeIpRipAuthenticationKeyChain976(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationStr974(d []interface{}) edpt.InterfaceVeIpRipAuthenticationStr974 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthenticationStr974
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationMode975(d []interface{}) edpt.InterfaceVeIpRipAuthenticationMode975 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthenticationMode975
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationKeyChain976(d []interface{}) edpt.InterfaceVeIpRipAuthenticationKeyChain976 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthenticationKeyChain976
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipSendCfg977(d []interface{}) edpt.InterfaceVeIpRipSendCfg977 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipSendCfg977
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipReceiveCfg978(d []interface{}) edpt.InterfaceVeIpRipReceiveCfg978 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipReceiveCfg978
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipSplitHorizonCfg979(d []interface{}) edpt.InterfaceVeIpRipSplitHorizonCfg979 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipSplitHorizonCfg979
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpOspf980(d []interface{}) edpt.InterfaceVeIpOspf980 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspf980
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceVeIpOspfOspfGlobal981(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceVeIpOspfOspfIpList988(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobal981(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobal981 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspfOspfGlobal981
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceVeIpOspfOspfGlobalAuthenticationCfg982(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceVeIpOspfOspfGlobalBfdCfg983(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceVeIpOspfOspfGlobalDatabaseFilterCfg984(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceVeIpOspfOspfGlobalMessageDigestCfg985(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceVeIpOspfOspfGlobalNetwork987(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalAuthenticationCfg982(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalAuthenticationCfg982 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspfOspfGlobalAuthenticationCfg982
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalBfdCfg983(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalBfdCfg983 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspfOspfGlobalBfdCfg983
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalDatabaseFilterCfg984(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg984 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg984
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceVeIpOspfOspfGlobalMessageDigestCfg985(d []interface{}) []edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfg985 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfg985, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfg985
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5986(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5986(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5986 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5986
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceVeIpOspfOspfGlobalNetwork987(d []interface{}) edpt.InterfaceVeIpOspfOspfGlobalNetwork987 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpOspfOspfGlobalNetwork987
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

func getSliceInterfaceVeIpOspfOspfIpList988(d []interface{}) []edpt.InterfaceVeIpOspfOspfIpList988 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpOspfOspfIpList988, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpOspfOspfIpList988
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceVeIpOspfOspfIpListMessageDigestCfg989(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpOspfOspfIpListMessageDigestCfg989(d []interface{}) []edpt.InterfaceVeIpOspfOspfIpListMessageDigestCfg989 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpOspfOspfIpListMessageDigestCfg989, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpOspfOspfIpListMessageDigestCfg989
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6990(d []interface{}) edpt.InterfaceVeIpv6990 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6990
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceInterfaceVeIpv6AddressList991(in["address_list"].([]interface{}))
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.V6AclName = in["v6_acl_name"].(string)
		ret.Inbound = in["inbound"].(int)
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.RouterAdver = getObjectInterfaceVeIpv6RouterAdver992(in["router_adver"].([]interface{}))
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceVeIpv6StatefulFirewall994(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceVeIpv6Router995(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceVeIpv6Rip1000(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceVeIpv6Ospf1002(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeIpv6AddressList991(d []interface{}) []edpt.InterfaceVeIpv6AddressList991 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6AddressList991, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6AddressList991
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterAdver992(d []interface{}) edpt.InterfaceVeIpv6RouterAdver992 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6RouterAdver992
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
		ret.AdverMtuDisable = in["adver_mtu_disable"].(int)
		ret.AdverMtu = in["adver_mtu"].(int)
		ret.PrefixList = getSliceInterfaceVeIpv6RouterAdverPrefixList993(in["prefix_list"].([]interface{}))
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

func getSliceInterfaceVeIpv6RouterAdverPrefixList993(d []interface{}) []edpt.InterfaceVeIpv6RouterAdverPrefixList993 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6RouterAdverPrefixList993, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6RouterAdverPrefixList993
		oi.Prefix = in["prefix"].(string)
		oi.NotAutonomous = in["not_autonomous"].(int)
		oi.NotOnLink = in["not_on_link"].(int)
		oi.PreferredLifetime = in["preferred_lifetime"].(int)
		oi.ValidLifetime = in["valid_lifetime"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6StatefulFirewall994(d []interface{}) edpt.InterfaceVeIpv6StatefulFirewall994 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6StatefulFirewall994
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

func getObjectInterfaceVeIpv6Router995(d []interface{}) edpt.InterfaceVeIpv6Router995 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6Router995
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceVeIpv6RouterRipng996(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceVeIpv6RouterOspf997(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceVeIpv6RouterIsis999(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterRipng996(d []interface{}) edpt.InterfaceVeIpv6RouterRipng996 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6RouterRipng996
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterOspf997(d []interface{}) edpt.InterfaceVeIpv6RouterOspf997 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6RouterOspf997
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceVeIpv6RouterOspfAreaList998(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceVeIpv6RouterOspfAreaList998(d []interface{}) []edpt.InterfaceVeIpv6RouterOspfAreaList998 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6RouterOspfAreaList998, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6RouterOspfAreaList998
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIpv6RouterIsis999(d []interface{}) edpt.InterfaceVeIpv6RouterIsis999 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6RouterIsis999
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpv6Rip1000(d []interface{}) edpt.InterfaceVeIpv6Rip1000 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6Rip1000
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceVeIpv6RipSplitHorizonCfg1001(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIpv6RipSplitHorizonCfg1001(d []interface{}) edpt.InterfaceVeIpv6RipSplitHorizonCfg1001 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6RipSplitHorizonCfg1001
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpv6Ospf1002(d []interface{}) edpt.InterfaceVeIpv6Ospf1002 {

	count1 := len(d)
	var ret edpt.InterfaceVeIpv6Ospf1002
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceVeIpv6OspfNetworkList1003(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceVeIpv6OspfCostCfg1004(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceVeIpv6OspfDeadIntervalCfg1005(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceVeIpv6OspfHelloIntervalCfg1006(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceVeIpv6OspfMtuIgnoreCfg1007(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceVeIpv6OspfNeighborCfg1008(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceVeIpv6OspfPriorityCfg1009(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceVeIpv6OspfRetransmitIntervalCfg1010(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceVeIpv6OspfTransmitDelayCfg1011(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfNetworkList1003(d []interface{}) []edpt.InterfaceVeIpv6OspfNetworkList1003 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfNetworkList1003, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfNetworkList1003
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfCostCfg1004(d []interface{}) []edpt.InterfaceVeIpv6OspfCostCfg1004 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfCostCfg1004, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfCostCfg1004
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfDeadIntervalCfg1005(d []interface{}) []edpt.InterfaceVeIpv6OspfDeadIntervalCfg1005 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfDeadIntervalCfg1005, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfDeadIntervalCfg1005
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfHelloIntervalCfg1006(d []interface{}) []edpt.InterfaceVeIpv6OspfHelloIntervalCfg1006 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfHelloIntervalCfg1006, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfHelloIntervalCfg1006
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfMtuIgnoreCfg1007(d []interface{}) []edpt.InterfaceVeIpv6OspfMtuIgnoreCfg1007 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfMtuIgnoreCfg1007, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfMtuIgnoreCfg1007
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfNeighborCfg1008(d []interface{}) []edpt.InterfaceVeIpv6OspfNeighborCfg1008 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfNeighborCfg1008, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfNeighborCfg1008
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfPriorityCfg1009(d []interface{}) []edpt.InterfaceVeIpv6OspfPriorityCfg1009 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfPriorityCfg1009, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfPriorityCfg1009
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfRetransmitIntervalCfg1010(d []interface{}) []edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg1010 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg1010, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg1010
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfTransmitDelayCfg1011(d []interface{}) []edpt.InterfaceVeIpv6OspfTransmitDelayCfg1011 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfTransmitDelayCfg1011, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfTransmitDelayCfg1011
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIsis1012(d []interface{}) edpt.InterfaceVeIsis1012 {

	count1 := len(d)
	var ret edpt.InterfaceVeIsis1012
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceVeIsisAuthentication1013(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceVeIsisBfdCfg1017(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceVeIsisCsnpIntervalList1018(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceVeIsisHelloIntervalList1019(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceVeIsisHelloIntervalMinimalList1020(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceVeIsisHelloMultiplierList1021(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceVeIsisMeshGroup1022(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceVeIsisMetricList1023(in["metric_list"].([]interface{}))
		ret.Network = in["network"].(string)
		ret.PasswordList = getSliceInterfaceVeIsisPasswordList1024(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceVeIsisPriorityList1025(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceVeIsisWideMetricList1026(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeIsisAuthentication1013(d []interface{}) edpt.InterfaceVeIsisAuthentication1013 {

	count1 := len(d)
	var ret edpt.InterfaceVeIsisAuthentication1013
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceVeIsisAuthenticationSendOnlyList1014(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceVeIsisAuthenticationModeList1015(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceVeIsisAuthenticationKeyChainList1016(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeIsisAuthenticationSendOnlyList1014(d []interface{}) []edpt.InterfaceVeIsisAuthenticationSendOnlyList1014 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisAuthenticationSendOnlyList1014, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisAuthenticationSendOnlyList1014
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisAuthenticationModeList1015(d []interface{}) []edpt.InterfaceVeIsisAuthenticationModeList1015 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisAuthenticationModeList1015, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisAuthenticationModeList1015
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisAuthenticationKeyChainList1016(d []interface{}) []edpt.InterfaceVeIsisAuthenticationKeyChainList1016 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisAuthenticationKeyChainList1016, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisAuthenticationKeyChainList1016
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIsisBfdCfg1017(d []interface{}) edpt.InterfaceVeIsisBfdCfg1017 {

	count1 := len(d)
	var ret edpt.InterfaceVeIsisBfdCfg1017
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceVeIsisCsnpIntervalList1018(d []interface{}) []edpt.InterfaceVeIsisCsnpIntervalList1018 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisCsnpIntervalList1018, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisCsnpIntervalList1018
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisHelloIntervalList1019(d []interface{}) []edpt.InterfaceVeIsisHelloIntervalList1019 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisHelloIntervalList1019, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisHelloIntervalList1019
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisHelloIntervalMinimalList1020(d []interface{}) []edpt.InterfaceVeIsisHelloIntervalMinimalList1020 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisHelloIntervalMinimalList1020, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisHelloIntervalMinimalList1020
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisHelloMultiplierList1021(d []interface{}) []edpt.InterfaceVeIsisHelloMultiplierList1021 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisHelloMultiplierList1021, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisHelloMultiplierList1021
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeIsisMeshGroup1022(d []interface{}) edpt.InterfaceVeIsisMeshGroup1022 {

	count1 := len(d)
	var ret edpt.InterfaceVeIsisMeshGroup1022
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceVeIsisMetricList1023(d []interface{}) []edpt.InterfaceVeIsisMetricList1023 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisMetricList1023, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisMetricList1023
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisPasswordList1024(d []interface{}) []edpt.InterfaceVeIsisPasswordList1024 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisPasswordList1024, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisPasswordList1024
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisPriorityList1025(d []interface{}) []edpt.InterfaceVeIsisPriorityList1025 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisPriorityList1025, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisPriorityList1025
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIsisWideMetricList1026(d []interface{}) []edpt.InterfaceVeIsisWideMetricList1026 {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIsisWideMetricList1026, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIsisWideMetricList1026
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceVeLw4o61027(d []interface{}) edpt.InterfaceVeLw4o61027 {

	count1 := len(d)
	var ret edpt.InterfaceVeLw4o61027
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceVeMap1028(d []interface{}) edpt.InterfaceVeMap1028 {

	count1 := len(d)
	var ret edpt.InterfaceVeMap1028
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

func getObjectInterfaceVeNptv61029(d []interface{}) edpt.InterfaceVeNptv61029 {

	count1 := len(d)
	var ret edpt.InterfaceVeNptv61029
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainList = getSliceInterfaceVeNptv6DomainList(in["domain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceVeNptv6DomainList(d []interface{}) []edpt.InterfaceVeNptv6DomainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeNptv6DomainList, 0, count1)
	for _, item := range d {
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

	count1 := len(d)
	ret := make([]edpt.InterfaceVeSamplingEnable, 0, count1)
	for _, item := range d {
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
	ret.Inst.Bfd = getObjectInterfaceVeBfd962(d.Get("bfd").([]interface{}))
	ret.Inst.Ddos = getObjectInterfaceVeDdos965(d.Get("ddos").([]interface{}))
	ret.Inst.GamingProtocolCompliance = d.Get("gaming_protocol_compliance").(int)
	ret.Inst.IcmpRateLimit = getObjectInterfaceVeIcmpRateLimit(d.Get("icmp_rate_limit").([]interface{}))
	ret.Inst.Icmpv6RateLimit = getObjectInterfaceVeIcmpv6RateLimit(d.Get("icmpv6_rate_limit").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceVeIp966(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectInterfaceVeIpv6990(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceVeIsis1012(d.Get("isis").([]interface{}))
	ret.Inst.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	ret.Inst.Lw4o6 = getObjectInterfaceVeLw4o61027(d.Get("lw_4o6").([]interface{}))
	ret.Inst.Map = getObjectInterfaceVeMap1028(d.Get("map").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Nptv6 = getObjectInterfaceVeNptv61029(d.Get("nptv6").([]interface{}))
	ret.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	ret.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	ret.Inst.SamplingEnable = getSliceInterfaceVeSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.TrapSource = d.Get("trap_source").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
