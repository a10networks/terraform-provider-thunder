package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet`: Ethernet interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetCreate,
		UpdateContext: resourceInterfaceEthernetUpdate,
		ReadContext:   resourceInterfaceEthernetRead,
		DeleteContext: resourceInterfaceEthernetDelete,

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
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable; 'disable': Disable;",
			},
			"auto_neg_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable auto-negotiation",
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
			"cpu_process": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All Packets to this port are processed by CPU",
			},
			"cpu_process_dir": {
				Type: schema.TypeString, Optional: true, Description: "'primary': Primary board; 'blade': blade board; 'hash-dip': Hash based on the Destination IP; 'hash-sip': Hash based on the Source IP; 'hash-dmac': Hash based on the Destination MAC; 'hash-smac': Hash based on the Source MAC;",
			},
			"ddos": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DDoS outside (untrusted) interface",
						},
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DDoS inside (trusted) interface",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"duplexity": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'Full': Full; 'Half': Half; 'auto': auto;",
			},
			"fec_forced_off": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "turn off the FEC",
			},
			"fec_forced_on": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "turn on the FEC",
			},
			"flow_control": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 802.3x flow control on full duplex port",
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
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet before sending out",
						},
						"syn_cookie": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure Enable SYN-cookie on the interface",
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
						"client": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client facing interface for IPv4/v6 traffic",
						},
						"server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server facing interface for IPv4/v6 traffic",
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
			"ipg_bit_time": {
				Type: schema.TypeInt, Optional: true, Default: 96, Description: "Set Inter-packet-gap interval in bit timing, default is 96",
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
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
						},
						"ipv6_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
						},
						"ttl_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet before sending out",
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
						"router_adver": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Router Advertisements on this interface; 'disable': Disable Router Advertisements on this interface;",
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
									"default_lifetime": {
										Type: schema.TypeInt, Optional: true, Default: 1800, Description: "Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))",
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"lldp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rt_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp enable/disable",
									},
									"rx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp rx",
									},
									"tx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp tx",
									},
								},
							},
						},
						"notification_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp notification configuration",
									},
									"notif_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp notification enable",
									},
								},
							},
						},
						"tx_dot1_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tx_dot1_tlvs": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration",
									},
									"link_aggregation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface link aggregation information",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface vlan information",
									},
								},
							},
						},
						"tx_tlvs_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tx_tlvs": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp tx TLVs configuration",
									},
									"exclude": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure which TLVs excluded. All basic TLVs will be included by default",
									},
									"management_address": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp management address",
									},
									"port_description": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp port description",
									},
									"system_capabilities": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system capabilities",
									},
									"system_description": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system description",
									},
									"system_name": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system name",
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
			"media_type_copper": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the media type to copper",
			},
			"monitor_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": {
							Type: schema.TypeString, Optional: true, Description: "'input': Incoming packets; 'output': Outgoing packets; 'both': Both incoming and outgoing packets;",
						},
						"mirror_index": {
							Type: schema.TypeInt, Optional: true, Description: "Mirror index",
						},
						"monitor_vlan": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN number",
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
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"ping_sweep_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enabl ping sweep detection; 'disable': Disable ping sweep detection(default);",
			},
			"port_breakout": {
				Type: schema.TypeString, Optional: true, Description: "'4x10G': Breakout 100G/40G ports into 4x10G ports; '4x25G': Breakout 100G ports into 4x25G ports; '2x50G': Breakout 100G ports into 2x50G ports;",
			},
			"port_scan_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable port scan detection; 'disable': Disable port scan detection(default);",
			},
			"remove_vlan_tag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove the vlan tag for egressing packets",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets_input': Input packets; 'bytes_input': Input bytes; 'received_broadcasts': Received broadcasts; 'received_multicasts': Received multicasts; 'received_unicasts': Received unicasts; 'input_errors': Input errors; 'crc': CRC; 'frame': Frames; 'runts': Runts; 'giants': Giants; 'packets_output': Output packets; 'bytes_output': Output bytes; 'transmitted_broadcasts': Transmitted broadcasts; 'transmitted_multicasts': Transmitted multicasts; 'transmitted_unicasts': Transmitted unicasts; 'output_errors': Output errors; 'collisions': Collisions; 'giants_output': Output Giants; 'rate_pkt_sent': Packet sent rate packets/sec; 'rate_byte_sent': Byte sent rate bits/sec; 'rate_pkt_rcvd': Packet received rate packets/sec; 'rate_byte_rcvd': Byte received rate bits/sec; 'load_interval': Load Interval; 'drops': Drops; 'input_utilization': Input Utilization; 'output_utilization': Output Utilization;",
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
			"speed": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'10': 10; '100': 100; '1000': 1000; 'auto': auto;",
			},
			"speed_forced_10g": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "force the speed to be 10G on 25G link",
			},
			"speed_forced_1g": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "force the speed to be 1G on 25G link",
			},
			"speed_forced_40g": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "force the speed to be 40G on 100G link",
			},
			"traffic_distribution_mode": {
				Type: schema.TypeString, Optional: true, Description: "'sip': sip; 'dip': dip; 'primary': primary; 'blade': blade; 'l4-src-port': l4-src-port; 'l4-dst-port': l4-dst-port;",
			},
			"trap_source": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The trap source",
			},
			"trunk_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk_number": {
							Type: schema.TypeInt, Required: true, Description: "Trunk Number",
						},
						"type": {
							Type: schema.TypeString, Optional: true, Default: "static", Description: "'static': Static (default); 'lacp': lacp; 'lacp-udld': lacp-udld;",
						},
						"admin_key": {
							Type: schema.TypeInt, Optional: true, Description: "LACP admin key (Admin key value)",
						},
						"port_priority": {
							Type: schema.TypeInt, Optional: true, Description: "Set LACP priority for a port (LACP port priority)",
						},
						"udld_timeout_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fast": {
										Type: schema.TypeInt, Optional: true, Default: 1000, Description: "fast timeout in unit of milli-seconds(default 1000)",
									},
									"slow": {
										Type: schema.TypeInt, Optional: true, Description: "slow timeout in unit of seconds",
									},
								},
							},
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Default: "active", Description: "'active': enable initiation of LACP negotiation on a port(default); 'passive': disable initiation of LACP negotiation on a port;",
						},
						"timeout": {
							Type: schema.TypeString, Optional: true, Default: "long", Description: "'long': Set LACP long timeout (default); 'short': Set LACP short timeout;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"update_l2_info": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Update and use received L2 information",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_wire": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark ethernet as a virtual wire interface",
			},
			"vlan_learning": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable VLAN learning; 'disable': Disable VLAN learning;",
			},
		},
	}
}
func resourceInterfaceEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceEthernetAccessList(d []interface{}) edpt.InterfaceEthernetAccessList {

	count1 := len(d)
	var ret edpt.InterfaceEthernetAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetBfd484(d []interface{}) edpt.InterfaceEthernetBfd484 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetBfd484
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceEthernetBfdAuthentication485(in["authentication"].([]interface{}))
		ret.Echo = in["echo"].(int)
		ret.Demand = in["demand"].(int)
		ret.IntervalCfg = getObjectInterfaceEthernetBfdIntervalCfg486(in["interval_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetBfdAuthentication485(d []interface{}) edpt.InterfaceEthernetBfdAuthentication485 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetBfdAuthentication485
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyId = in["key_id"].(int)
		ret.Method = in["method"].(string)
		ret.Password = in["password"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceEthernetBfdIntervalCfg486(d []interface{}) edpt.InterfaceEthernetBfdIntervalCfg486 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetBfdIntervalCfg486
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.MinRx = in["min_rx"].(int)
		ret.Multiplier = in["multiplier"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetDdos487(d []interface{}) edpt.InterfaceEthernetDdos487 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetDdos487
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIcmpRateLimit(d []interface{}) edpt.InterfaceEthernetIcmpRateLimit {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIcmpRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Normal = in["normal"].(int)
		ret.Lockup = in["lockup"].(int)
		ret.LockupPeriod = in["lockup_period"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetIcmpv6RateLimit(d []interface{}) edpt.InterfaceEthernetIcmpv6RateLimit {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIcmpv6RateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NormalV6 = in["normal_v6"].(int)
		ret.LockupV6 = in["lockup_v6"].(int)
		ret.LockupPeriodV6 = in["lockup_period_v6"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetIp488(d []interface{}) edpt.InterfaceEthernetIp488 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIp488
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.AddressList = getSliceInterfaceEthernetIpAddressList489(in["address_list"].([]interface{}))
		ret.AllowPromiscuousVip = in["allow_promiscuous_vip"].(int)
		ret.CacheSpoofingPort = in["cache_spoofing_port"].(int)
		ret.HelperAddressList = getSliceInterfaceEthernetIpHelperAddressList490(in["helper_address_list"].([]interface{}))
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.SynCookie = in["syn_cookie"].(int)
		ret.SlbPartitionRedirect = in["slb_partition_redirect"].(int)
		ret.GenerateMembershipQuery = in["generate_membership_query"].(int)
		ret.QueryInterval = in["query_interval"].(int)
		ret.MaxRespTime = in["max_resp_time"].(int)
		ret.Client = in["client"].(int)
		ret.Server = in["server"].(int)
		ret.Unnumbered = in["unnumbered"].(int)
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceEthernetIpStatefulFirewall491(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceEthernetIpRouter492(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceEthernetIpRip494(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceEthernetIpOspf502(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceEthernetIpAddressList489(d []interface{}) []edpt.InterfaceEthernetIpAddressList489 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpAddressList489, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpAddressList489
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpHelperAddressList490(d []interface{}) []edpt.InterfaceEthernetIpHelperAddressList490 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpHelperAddressList490, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpHelperAddressList490
		oi.HelperAddress = in["helper_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpStatefulFirewall491(d []interface{}) edpt.InterfaceEthernetIpStatefulFirewall491 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpStatefulFirewall491
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

func getObjectInterfaceEthernetIpRouter492(d []interface{}) edpt.InterfaceEthernetIpRouter492 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRouter492
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceEthernetIpRouterIsis493(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpRouterIsis493(d []interface{}) edpt.InterfaceEthernetIpRouterIsis493 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRouterIsis493
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpRip494(d []interface{}) edpt.InterfaceEthernetIpRip494 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRip494
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceEthernetIpRipAuthentication495(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceEthernetIpRipSendCfg499(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceEthernetIpRipReceiveCfg500(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceEthernetIpRipSplitHorizonCfg501(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthentication495(d []interface{}) edpt.InterfaceEthernetIpRipAuthentication495 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthentication495
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceEthernetIpRipAuthenticationStr496(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceEthernetIpRipAuthenticationMode497(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceEthernetIpRipAuthenticationKeyChain498(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthenticationStr496(d []interface{}) edpt.InterfaceEthernetIpRipAuthenticationStr496 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthenticationStr496
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthenticationMode497(d []interface{}) edpt.InterfaceEthernetIpRipAuthenticationMode497 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthenticationMode497
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthenticationKeyChain498(d []interface{}) edpt.InterfaceEthernetIpRipAuthenticationKeyChain498 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthenticationKeyChain498
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipSendCfg499(d []interface{}) edpt.InterfaceEthernetIpRipSendCfg499 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipSendCfg499
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipReceiveCfg500(d []interface{}) edpt.InterfaceEthernetIpRipReceiveCfg500 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipReceiveCfg500
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipSplitHorizonCfg501(d []interface{}) edpt.InterfaceEthernetIpRipSplitHorizonCfg501 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipSplitHorizonCfg501
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspf502(d []interface{}) edpt.InterfaceEthernetIpOspf502 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspf502
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceEthernetIpOspfOspfGlobal503(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceEthernetIpOspfOspfIpList510(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobal503(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobal503 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobal503
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceEthernetIpOspfOspfGlobalAuthenticationCfg504(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceEthernetIpOspfOspfGlobalBfdCfg505(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg506(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceEthernetIpOspfOspfGlobalNetwork509(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalAuthenticationCfg504(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg504 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg504
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalBfdCfg505(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalBfdCfg505 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalBfdCfg505
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg506(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg506 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg506
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5508(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5508(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5508 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5508
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalNetwork509(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalNetwork509 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalNetwork509
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

func getSliceInterfaceEthernetIpOspfOspfIpList510(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfIpList510 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfIpList510, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfIpList510
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceEthernetIpOspfOspfIpListMessageDigestCfg511(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpOspfOspfIpListMessageDigestCfg511(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfIpListMessageDigestCfg511 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfIpListMessageDigestCfg511, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfIpListMessageDigestCfg511
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6512(d []interface{}) edpt.InterfaceEthernetIpv6512 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6512
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceInterfaceEthernetIpv6AddressList513(in["address_list"].([]interface{}))
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
		ret.Ipv6Enable = in["ipv6_enable"].(int)
		ret.TtlIgnore = in["ttl_ignore"].(int)
		ret.AccessListCfg = getObjectInterfaceEthernetIpv6AccessListCfg514(in["access_list_cfg"].([]interface{}))
		ret.RouterAdver = getObjectInterfaceEthernetIpv6RouterAdver515(in["router_adver"].([]interface{}))
		//omit uuid
		ret.StatefulFirewall = getObjectInterfaceEthernetIpv6StatefulFirewall517(in["stateful_firewall"].([]interface{}))
		ret.Router = getObjectInterfaceEthernetIpv6Router518(in["router"].([]interface{}))
		ret.Rip = getObjectInterfaceEthernetIpv6Rip523(in["rip"].([]interface{}))
		ret.Ospf = getObjectInterfaceEthernetIpv6Ospf525(in["ospf"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceEthernetIpv6AddressList513(d []interface{}) []edpt.InterfaceEthernetIpv6AddressList513 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6AddressList513, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6AddressList513
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6AccessListCfg514(d []interface{}) edpt.InterfaceEthernetIpv6AccessListCfg514 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6AccessListCfg514
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V6AclName = in["v6_acl_name"].(string)
		ret.Inbound = in["inbound"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterAdver515(d []interface{}) edpt.InterfaceEthernetIpv6RouterAdver515 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterAdver515
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.HopLimit = in["hop_limit"].(int)
		ret.MaxInterval = in["max_interval"].(int)
		ret.MinInterval = in["min_interval"].(int)
		ret.DefaultLifetime = in["default_lifetime"].(int)
		ret.RateLimit = in["rate_limit"].(int)
		ret.ReachableTime = in["reachable_time"].(int)
		ret.RetransmitTimer = in["retransmit_timer"].(int)
		ret.AdverMtuDisable = in["adver_mtu_disable"].(int)
		ret.AdverMtu = in["adver_mtu"].(int)
		ret.PrefixList = getSliceInterfaceEthernetIpv6RouterAdverPrefixList516(in["prefix_list"].([]interface{}))
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

func getSliceInterfaceEthernetIpv6RouterAdverPrefixList516(d []interface{}) []edpt.InterfaceEthernetIpv6RouterAdverPrefixList516 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6RouterAdverPrefixList516, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6RouterAdverPrefixList516
		oi.Prefix = in["prefix"].(string)
		oi.NotAutonomous = in["not_autonomous"].(int)
		oi.NotOnLink = in["not_on_link"].(int)
		oi.PreferredLifetime = in["preferred_lifetime"].(int)
		oi.ValidLifetime = in["valid_lifetime"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6StatefulFirewall517(d []interface{}) edpt.InterfaceEthernetIpv6StatefulFirewall517 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6StatefulFirewall517
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

func getObjectInterfaceEthernetIpv6Router518(d []interface{}) edpt.InterfaceEthernetIpv6Router518 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6Router518
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceEthernetIpv6RouterRipng519(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceEthernetIpv6RouterOspf520(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceEthernetIpv6RouterIsis522(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterRipng519(d []interface{}) edpt.InterfaceEthernetIpv6RouterRipng519 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterRipng519
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterOspf520(d []interface{}) edpt.InterfaceEthernetIpv6RouterOspf520 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterOspf520
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceEthernetIpv6RouterOspfAreaList521(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceEthernetIpv6RouterOspfAreaList521(d []interface{}) []edpt.InterfaceEthernetIpv6RouterOspfAreaList521 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6RouterOspfAreaList521, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6RouterOspfAreaList521
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RouterIsis522(d []interface{}) edpt.InterfaceEthernetIpv6RouterIsis522 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RouterIsis522
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpv6Rip523(d []interface{}) edpt.InterfaceEthernetIpv6Rip523 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6Rip523
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceEthernetIpv6RipSplitHorizonCfg524(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpv6RipSplitHorizonCfg524(d []interface{}) edpt.InterfaceEthernetIpv6RipSplitHorizonCfg524 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6RipSplitHorizonCfg524
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpv6Ospf525(d []interface{}) edpt.InterfaceEthernetIpv6Ospf525 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpv6Ospf525
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceEthernetIpv6OspfNetworkList526(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceEthernetIpv6OspfCostCfg527(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceEthernetIpv6OspfDeadIntervalCfg528(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceEthernetIpv6OspfHelloIntervalCfg529(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceEthernetIpv6OspfMtuIgnoreCfg530(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceEthernetIpv6OspfNeighborCfg531(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceEthernetIpv6OspfPriorityCfg532(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceEthernetIpv6OspfRetransmitIntervalCfg533(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceEthernetIpv6OspfTransmitDelayCfg534(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfNetworkList526(d []interface{}) []edpt.InterfaceEthernetIpv6OspfNetworkList526 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfNetworkList526, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfNetworkList526
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfCostCfg527(d []interface{}) []edpt.InterfaceEthernetIpv6OspfCostCfg527 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfCostCfg527, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfCostCfg527
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfDeadIntervalCfg528(d []interface{}) []edpt.InterfaceEthernetIpv6OspfDeadIntervalCfg528 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfDeadIntervalCfg528, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfDeadIntervalCfg528
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfHelloIntervalCfg529(d []interface{}) []edpt.InterfaceEthernetIpv6OspfHelloIntervalCfg529 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfHelloIntervalCfg529, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfHelloIntervalCfg529
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfMtuIgnoreCfg530(d []interface{}) []edpt.InterfaceEthernetIpv6OspfMtuIgnoreCfg530 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfMtuIgnoreCfg530, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfMtuIgnoreCfg530
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfNeighborCfg531(d []interface{}) []edpt.InterfaceEthernetIpv6OspfNeighborCfg531 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfNeighborCfg531, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfNeighborCfg531
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfPriorityCfg532(d []interface{}) []edpt.InterfaceEthernetIpv6OspfPriorityCfg532 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfPriorityCfg532, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfPriorityCfg532
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfRetransmitIntervalCfg533(d []interface{}) []edpt.InterfaceEthernetIpv6OspfRetransmitIntervalCfg533 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfRetransmitIntervalCfg533, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfRetransmitIntervalCfg533
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpv6OspfTransmitDelayCfg534(d []interface{}) []edpt.InterfaceEthernetIpv6OspfTransmitDelayCfg534 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpv6OspfTransmitDelayCfg534, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpv6OspfTransmitDelayCfg534
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIsis535(d []interface{}) edpt.InterfaceEthernetIsis535 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsis535
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceEthernetIsisAuthentication536(in["authentication"].([]interface{}))
		ret.BfdCfg = getObjectInterfaceEthernetIsisBfdCfg540(in["bfd_cfg"].([]interface{}))
		ret.CircuitType = in["circuit_type"].(string)
		ret.CsnpIntervalList = getSliceInterfaceEthernetIsisCsnpIntervalList541(in["csnp_interval_list"].([]interface{}))
		ret.Padding = in["padding"].(int)
		ret.HelloIntervalList = getSliceInterfaceEthernetIsisHelloIntervalList542(in["hello_interval_list"].([]interface{}))
		ret.HelloIntervalMinimalList = getSliceInterfaceEthernetIsisHelloIntervalMinimalList543(in["hello_interval_minimal_list"].([]interface{}))
		ret.HelloMultiplierList = getSliceInterfaceEthernetIsisHelloMultiplierList544(in["hello_multiplier_list"].([]interface{}))
		ret.LspInterval = in["lsp_interval"].(int)
		ret.MeshGroup = getObjectInterfaceEthernetIsisMeshGroup545(in["mesh_group"].([]interface{}))
		ret.MetricList = getSliceInterfaceEthernetIsisMetricList546(in["metric_list"].([]interface{}))
		ret.Network = in["network"].(string)
		ret.PasswordList = getSliceInterfaceEthernetIsisPasswordList547(in["password_list"].([]interface{}))
		ret.PriorityList = getSliceInterfaceEthernetIsisPriorityList548(in["priority_list"].([]interface{}))
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.WideMetricList = getSliceInterfaceEthernetIsisWideMetricList549(in["wide_metric_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIsisAuthentication536(d []interface{}) edpt.InterfaceEthernetIsisAuthentication536 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsisAuthentication536
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceEthernetIsisAuthenticationSendOnlyList537(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceEthernetIsisAuthenticationModeList538(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceEthernetIsisAuthenticationKeyChainList539(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceEthernetIsisAuthenticationSendOnlyList537(d []interface{}) []edpt.InterfaceEthernetIsisAuthenticationSendOnlyList537 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisAuthenticationSendOnlyList537, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisAuthenticationSendOnlyList537
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisAuthenticationModeList538(d []interface{}) []edpt.InterfaceEthernetIsisAuthenticationModeList538 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisAuthenticationModeList538, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisAuthenticationModeList538
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisAuthenticationKeyChainList539(d []interface{}) []edpt.InterfaceEthernetIsisAuthenticationKeyChainList539 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisAuthenticationKeyChainList539, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisAuthenticationKeyChainList539
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIsisBfdCfg540(d []interface{}) edpt.InterfaceEthernetIsisBfdCfg540 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsisBfdCfg540
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIsisCsnpIntervalList541(d []interface{}) []edpt.InterfaceEthernetIsisCsnpIntervalList541 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisCsnpIntervalList541, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisCsnpIntervalList541
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisHelloIntervalList542(d []interface{}) []edpt.InterfaceEthernetIsisHelloIntervalList542 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisHelloIntervalList542, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisHelloIntervalList542
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisHelloIntervalMinimalList543(d []interface{}) []edpt.InterfaceEthernetIsisHelloIntervalMinimalList543 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisHelloIntervalMinimalList543, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisHelloIntervalMinimalList543
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisHelloMultiplierList544(d []interface{}) []edpt.InterfaceEthernetIsisHelloMultiplierList544 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisHelloMultiplierList544, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisHelloMultiplierList544
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIsisMeshGroup545(d []interface{}) edpt.InterfaceEthernetIsisMeshGroup545 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsisMeshGroup545
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIsisMetricList546(d []interface{}) []edpt.InterfaceEthernetIsisMetricList546 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisMetricList546, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisMetricList546
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisPasswordList547(d []interface{}) []edpt.InterfaceEthernetIsisPasswordList547 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisPasswordList547, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisPasswordList547
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisPriorityList548(d []interface{}) []edpt.InterfaceEthernetIsisPriorityList548 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisPriorityList548, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisPriorityList548
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisWideMetricList549(d []interface{}) []edpt.InterfaceEthernetIsisWideMetricList549 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisWideMetricList549, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisWideMetricList549
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetLldp550(d []interface{}) edpt.InterfaceEthernetLldp550 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldp550
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableCfg = getObjectInterfaceEthernetLldpEnableCfg551(in["enable_cfg"].([]interface{}))
		ret.NotificationCfg = getObjectInterfaceEthernetLldpNotificationCfg552(in["notification_cfg"].([]interface{}))
		ret.TxDot1Cfg = getObjectInterfaceEthernetLldpTxDot1Cfg553(in["tx_dot1_cfg"].([]interface{}))
		ret.TxTlvsCfg = getObjectInterfaceEthernetLldpTxTlvsCfg554(in["tx_tlvs_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetLldpEnableCfg551(d []interface{}) edpt.InterfaceEthernetLldpEnableCfg551 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpEnableCfg551
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RtEnable = in["rt_enable"].(int)
		ret.Rx = in["rx"].(int)
		ret.Tx = in["tx"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLldpNotificationCfg552(d []interface{}) edpt.InterfaceEthernetLldpNotificationCfg552 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpNotificationCfg552
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = in["notification"].(int)
		ret.NotifEnable = in["notif_enable"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLldpTxDot1Cfg553(d []interface{}) edpt.InterfaceEthernetLldpTxDot1Cfg553 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpTxDot1Cfg553
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxDot1Tlvs = in["tx_dot1_tlvs"].(int)
		ret.LinkAggregation = in["link_aggregation"].(int)
		ret.Vlan = in["vlan"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLldpTxTlvsCfg554(d []interface{}) edpt.InterfaceEthernetLldpTxTlvsCfg554 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLldpTxTlvsCfg554
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxTlvs = in["tx_tlvs"].(int)
		ret.Exclude = in["exclude"].(int)
		ret.ManagementAddress = in["management_address"].(int)
		ret.PortDescription = in["port_description"].(int)
		ret.SystemCapabilities = in["system_capabilities"].(int)
		ret.SystemDescription = in["system_description"].(int)
		ret.SystemName = in["system_name"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetLw4o6555(d []interface{}) edpt.InterfaceEthernetLw4o6555 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetLw4o6555
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Outside = in["outside"].(int)
		ret.Inside = in["inside"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetMap556(d []interface{}) edpt.InterfaceEthernetMap556 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetMap556
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

func getSliceInterfaceEthernetMonitorList(d []interface{}) []edpt.InterfaceEthernetMonitorList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetMonitorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetMonitorList
		oi.Monitor = in["monitor"].(string)
		oi.MirrorIndex = in["mirror_index"].(int)
		oi.MonitorVlan = in["monitor_vlan"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetNptv6557(d []interface{}) edpt.InterfaceEthernetNptv6557 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetNptv6557
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainList = getSliceInterfaceEthernetNptv6DomainList(in["domain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceEthernetNptv6DomainList(d []interface{}) []edpt.InterfaceEthernetNptv6DomainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetNptv6DomainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetNptv6DomainList
		oi.DomainName = in["domain_name"].(string)
		oi.BindType = in["bind_type"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetSamplingEnable(d []interface{}) []edpt.InterfaceEthernetSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetSpanningTree558(d []interface{}) edpt.InterfaceEthernetSpanningTree558 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetSpanningTree558
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AutoEdge = in["auto_edge"].(int)
		ret.AdminEdge = in["admin_edge"].(int)
		ret.InstanceList = getSliceInterfaceEthernetSpanningTreeInstanceList559(in["instance_list"].([]interface{}))
		ret.PathCost = in["path_cost"].(int)
		//omit uuid
	}
	return ret
}

func getSliceInterfaceEthernetSpanningTreeInstanceList559(d []interface{}) []edpt.InterfaceEthernetSpanningTreeInstanceList559 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetSpanningTreeInstanceList559, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetSpanningTreeInstanceList559
		oi.InstanceStart = in["instance_start"].(int)
		oi.MstpPathCost = in["mstp_path_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetTrunkGroupList(d []interface{}) []edpt.InterfaceEthernetTrunkGroupList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetTrunkGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetTrunkGroupList
		oi.TrunkNumber = in["trunk_number"].(int)
		oi.Type = in["type"].(string)
		oi.AdminKey = in["admin_key"].(int)
		oi.PortPriority = in["port_priority"].(int)
		oi.UdldTimeoutCfg = getObjectInterfaceEthernetTrunkGroupListUdldTimeoutCfg(in["udld_timeout_cfg"].([]interface{}))
		oi.Mode = in["mode"].(string)
		oi.Timeout = in["timeout"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetTrunkGroupListUdldTimeoutCfg(d []interface{}) edpt.InterfaceEthernetTrunkGroupListUdldTimeoutCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetTrunkGroupListUdldTimeoutCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Fast = in["fast"].(int)
		ret.Slow = in["slow"].(int)
	}
	return ret
}

func dataToEndpointInterfaceEthernet(d *schema.ResourceData) edpt.InterfaceEthernet {
	var ret edpt.InterfaceEthernet
	ret.Inst.AccessList = getObjectInterfaceEthernetAccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AutoNegEnable = d.Get("auto_neg_enable").(int)
	ret.Inst.Bfd = getObjectInterfaceEthernetBfd484(d.Get("bfd").([]interface{}))
	ret.Inst.CpuProcess = d.Get("cpu_process").(int)
	ret.Inst.CpuProcessDir = d.Get("cpu_process_dir").(string)
	ret.Inst.Ddos = getObjectInterfaceEthernetDdos487(d.Get("ddos").([]interface{}))
	ret.Inst.Duplexity = d.Get("duplexity").(string)
	ret.Inst.FecForcedOff = d.Get("fec_forced_off").(int)
	ret.Inst.FecForcedOn = d.Get("fec_forced_on").(int)
	ret.Inst.FlowControl = d.Get("flow_control").(int)
	ret.Inst.GamingProtocolCompliance = d.Get("gaming_protocol_compliance").(int)
	ret.Inst.IcmpRateLimit = getObjectInterfaceEthernetIcmpRateLimit(d.Get("icmp_rate_limit").([]interface{}))
	ret.Inst.Icmpv6RateLimit = getObjectInterfaceEthernetIcmpv6RateLimit(d.Get("icmpv6_rate_limit").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(int)
	ret.Inst.Ip = getObjectInterfaceEthernetIp488(d.Get("ip").([]interface{}))
	ret.Inst.IpgBitTime = d.Get("ipg_bit_time").(int)
	ret.Inst.Ipv6 = getObjectInterfaceEthernetIpv6512(d.Get("ipv6").([]interface{}))
	ret.Inst.Isis = getObjectInterfaceEthernetIsis535(d.Get("isis").([]interface{}))
	ret.Inst.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	ret.Inst.Lldp = getObjectInterfaceEthernetLldp550(d.Get("lldp").([]interface{}))
	ret.Inst.LoadInterval = d.Get("load_interval").(int)
	ret.Inst.Lw4o6 = getObjectInterfaceEthernetLw4o6555(d.Get("lw_4o6").([]interface{}))
	ret.Inst.MacLearning = d.Get("mac_learning").(string)
	ret.Inst.Map = getObjectInterfaceEthernetMap556(d.Get("map").([]interface{}))
	ret.Inst.MediaTypeCopper = d.Get("media_type_copper").(int)
	ret.Inst.MonitorList = getSliceInterfaceEthernetMonitorList(d.Get("monitor_list").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Nptv6 = getObjectInterfaceEthernetNptv6557(d.Get("nptv6").([]interface{}))
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	ret.Inst.PortBreakout = d.Get("port_breakout").(string)
	ret.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	ret.Inst.RemoveVlanTag = d.Get("remove_vlan_tag").(int)
	ret.Inst.SamplingEnable = getSliceInterfaceEthernetSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SpanningTree = getObjectInterfaceEthernetSpanningTree558(d.Get("spanning_tree").([]interface{}))
	ret.Inst.Speed = d.Get("speed").(string)
	ret.Inst.SpeedForced10g = d.Get("speed_forced_10g").(int)
	ret.Inst.SpeedForced1g = d.Get("speed_forced_1g").(int)
	ret.Inst.SpeedForced40g = d.Get("speed_forced_40g").(int)
	ret.Inst.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)
	ret.Inst.TrapSource = d.Get("trap_source").(int)
	ret.Inst.TrunkGroupList = getSliceInterfaceEthernetTrunkGroupList(d.Get("trunk_group_list").([]interface{}))
	ret.Inst.UpdateL2Info = d.Get("update_l2_info").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VirtualWire = d.Get("virtual_wire").(int)
	ret.Inst.VlanLearning = d.Get("vlan_learning").(string)
	return ret
}
