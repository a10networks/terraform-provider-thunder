package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_isis`: Intermediate System - Intermediate System (IS-IS)\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIsisCreate,
		UpdateContext: resourceRouterIsisUpdate,
		ReadContext:   resourceRouterIsisRead,
		DeleteContext: resourceRouterIsisDelete,
		Schema: map[string]*schema.Schema{
			"address_family": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_information": {
										Type: schema.TypeString, Optional: true, Description: "'originate': Distribute a default route;",
									},
									"adjacency_check": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Check ISIS neighbor protocol support",
									},
									"distance": {
										Type: schema.TypeInt, Optional: true, Default: 115, Description: "ISIS Administrative Distance (Distance value)",
									},
									"multi_topology_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"multi_topology": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multi-topology mode",
												},
												"level": {
													Type: schema.TypeString, Optional: true, Description: "'level-1': Level-1 only; 'level-1-2': Level-1-2; 'level-2': Level-2 only;",
												},
												"transition": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept and generate both IS-IS IPv6 and Multi-topology IPV6 TLVs",
												},
												"level_transition": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept and generate both IS-IS IPv6 and Multi-topology IPV6 TLVs",
												},
											},
										},
									},
									"summary_prefix_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"prefix": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
												},
												"level": {
													Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': Summarize into level-1 area; 'level-1-2': Summarize into both area and sub-domain; 'level-2': Summarize into level-2 sub-domain;",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"redistribute": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"redist_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'ip-nat': IP NAT; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'nat64': NAT64 Prefix; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
															},
															"metric": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Metric for redistributed routes (IS-IS default metric)",
															},
															"metric_type": {
																Type: schema.TypeString, Optional: true, Default: "internal", Description: "'external': Set IS-IS External metric type; 'internal': Set IS-IS Internal metric type;",
															},
															"route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
															"level": {
																Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': IS-IS level-1 routes only; 'level-1-2': IS-IS level-1 and level-2 routes; 'level-2': IS-IS level-2 routes only;",
															},
														},
													},
												},
												"vip_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"vip_type": {
																Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
															},
															"vip_metric": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Metric for redistributed routes (IS-IS default metric)",
															},
															"vip_route_map": {
																Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
															},
															"vip_metric_type": {
																Type: schema.TypeString, Optional: true, Default: "internal", Description: "'external': Set IS-IS External metric type; 'internal': Set IS-IS Internal metric type;",
															},
															"vip_level": {
																Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': IS-IS level-1 routes only; 'level-1-2': IS-IS level-1 and level-2 routes; 'level-2': IS-IS level-2 routes only;",
															},
														},
													},
												},
												"isis": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"level_1_from": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"into_1": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"level_2": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-area routes into level-2",
																					},
																					"distribute_list": {
																						Type: schema.TypeString, Optional: true, Description: "Select routes (Access-list name)",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"level_2_from": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"into_2": {
																			Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"level_1": {
																						Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-area routes into level-2",
																					},
																					"distribute_list": {
																						Type: schema.TypeString, Optional: true, Description: "Select routes (Access-list name)",
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
			"adjacency_check": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Check ISIS neighbor protocol support",
			},
			"area_password_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type: schema.TypeString, Optional: true, Description: "Configure the authentication password for an area (Area password)",
						},
						"authenticate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"snp": {
										Type: schema.TypeString, Optional: true, Description: "'send-only': Send but do not check PDUs on receiving; 'validate': Send and check PDUs on receiving;",
									},
								},
							},
						},
					},
				},
			},
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
										Type: schema.TypeString, Optional: true, Description: "'md5': Authentication mode;",
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
			"bfd": {
				Type: schema.TypeString, Optional: true, Description: "'all-interfaces': Enable BFD on all interfaces;",
			},
			"default_information": {
				Type: schema.TypeString, Optional: true, Description: "'originate': Distribute a default route;",
			},
			"distance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": {
							Type: schema.TypeInt, Optional: true, Default: 115, Description: "ISIS Administrative Distance (Distance value)",
						},
						"system_id": {
							Type: schema.TypeString, Optional: true, Description: "System-ID in XXXX.XXXX.XXXX",
						},
						"acl": {
							Type: schema.TypeString, Optional: true, Description: "Access list name",
						},
					},
				},
			},
			"domain_password_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type: schema.TypeString, Optional: true, Description: "Set the authentication password for a routing domain (Routing domain password)",
						},
						"authenticate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"snp": {
										Type: schema.TypeString, Optional: true, Description: "'send-only': Send but do not check PDUs on receiving; 'validate': Send and check PDUs on receiving;",
									},
								},
							},
						},
					},
				},
			},
			"ha_standby_extra_cost": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extra_cost": {
							Type: schema.TypeInt, Optional: true, Description: "The extra cost value",
						},
						"group": {
							Type: schema.TypeInt, Optional: true, Description: "Group (Group ID)",
						},
					},
				},
			},
			"ignore_lsp_errors": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore LSPs with bad checksums",
			},
			"is_type": {
				Type: schema.TypeString, Optional: true, Default: "level-1", Description: "'level-1': Act as a station router only; 'level-1-2': Act as both a station router and an area router; 'level-2-only': Act as an area router only;",
			},
			"log_adjacency_changes_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "'detail': Log changes in adjacency state; 'disable': Disable logging;",
						},
					},
				},
			},
			"lsp_gen_interval_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Minimum interval in seconds",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Set interval for level 1 only; 'level-2': Set interval for level 2 only;",
						},
					},
				},
			},
			"lsp_refresh_interval": {
				Type: schema.TypeInt, Optional: true, Default: 900, Description: "Set LSP refresh interval (LSP refresh time in seconds)",
			},
			"max_lsp_lifetime": {
				Type: schema.TypeInt, Optional: true, Default: 1200, Description: "Set maximum LSP lifetime (Maximum LSP lifetime in seconds)",
			},
			"metric_style_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Default: "narrow", Description: "'narrow': Use old style of TLVs with narrow metric; 'wide': Use new style of TLVs to carry wider metric; 'transition': Send and accept both styles of TLVs during transition; 'narrow-transition': Send old style of TLVs with narrow metric with accepting both styles of TLVs; 'wide-transition': Send new style of TLVs to carry wider metric with accepting both styles of TLVs;",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Default: "level-1-2", Description: "'level-1': Level-1 only; 'level-1-2': Level-1-2; 'level-2': Level-2 only;",
						},
					},
				},
			},
			"net_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"net": {
							Type: schema.TypeString, Optional: true, Description: "A Network Entity Title for this process (XX.XXXX. ... .XXXX.XX  Network entity title (NET))",
						},
					},
				},
			},
			"passive_interface_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"lif": {
							Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
						},
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
					},
				},
			},
			"protocol_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol_topology": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Protocol Topology",
						},
					},
				},
			},
			"redistribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redist_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'ip-nat': IP NAT; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Metric for redistributed routes (IS-IS default metric)",
									},
									"metric_type": {
										Type: schema.TypeString, Optional: true, Default: "internal", Description: "'external': Set IS-IS External metric type; 'internal': Set IS-IS Internal metric type;",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': IS-IS level-1 routes only; 'level-1-2': IS-IS level-1 and level-2 routes; 'level-2': IS-IS level-2 routes only;",
									},
								},
							},
						},
						"vip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_type": {
										Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
									},
									"vip_metric": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Metric for redistributed routes (IS-IS default metric)",
									},
									"vip_route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
									"vip_metric_type": {
										Type: schema.TypeString, Optional: true, Default: "internal", Description: "'external': Set IS-IS External metric type; 'internal': Set IS-IS Internal metric type;",
									},
									"vip_level": {
										Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': IS-IS level-1 routes only; 'level-1-2': IS-IS level-1 and level-2 routes; 'level-2': IS-IS level-2 routes only;",
									},
								},
							},
						},
						"isis": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level_1_from": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"into_1": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"level_2": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-area routes into level-2",
															},
															"distribute_list": {
																Type: schema.TypeString, Optional: true, Description: "Select routes (Access-list name)",
															},
														},
													},
												},
											},
										},
									},
									"level_2_from": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"into_2": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"level_1": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-area routes into level-2",
															},
															"distribute_list": {
																Type: schema.TypeString, Optional: true, Description: "Select routes (Access-list name)",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"set_overload_bit_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set_overload_bit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Signal other touers not to use us in SPF",
						},
						"on_startup": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delay": {
										Type: schema.TypeInt, Optional: true, Description: "Time in seconds to advertise ourself as overloaded after reboot",
									},
									"wait_for_bgp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Let BGP decide when to unset the overload bit",
									},
								},
							},
						},
						"suppress_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "If overload-bit set, don't advertise IP prefixes learned from other protocols",
									},
									"interlevel": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "If overload-bit set, don't advertise IP prefixes learned from another ISIS level",
									},
								},
							},
						},
					},
				},
			},
			"spf_interval_exp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min": {
							Type: schema.TypeInt, Optional: true, Default: 500, Description: "Minimum Delay between receiving a change to SPF calculation in milliseconds",
						},
						"max": {
							Type: schema.TypeInt, Optional: true, Default: 50000, Description: "Maximum Delay between receiving a change to SPF calculation in milliseconds",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Set interval for level 1 only; 'level-2': Set interval for level 2 only;",
						},
					},
				},
			},
			"summary_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type: schema.TypeString, Optional: true, Description: "IP network prefix",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': Summarize into level-1 area; 'level-1-2': Summarize into both area and sub-domain; 'level-2': Summarize into level-2 sub-domain;",
						},
					},
				},
			},
			"tag": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "ISO routing area tag",
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

func resourceRouterIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterIsisAddressFamily(d []interface{}) edpt.RouterIsisAddressFamily {
	count := len(d)
	var ret edpt.RouterIsisAddressFamily
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6 = getObjectRouterIsisAddressFamilyIpv6(in["ipv6"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6(d []interface{}) edpt.RouterIsisAddressFamilyIpv6 {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultInformation = in["default_information"].(string)
		ret.AdjacencyCheck = in["adjacency_check"].(int)
		ret.Distance = in["distance"].(int)
		ret.MultiTopologyCfg = getObjectRouterIsisAddressFamilyIpv6MultiTopologyCfg(in["multi_topology_cfg"].([]interface{}))
		ret.SummaryPrefixList = getSliceRouterIsisAddressFamilyIpv6SummaryPrefixList(in["summary_prefix_list"].([]interface{}))
		//omit uuid
		ret.Redistribute = getObjectRouterIsisAddressFamilyIpv6Redistribute(in["redistribute"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6MultiTopologyCfg(d []interface{}) edpt.RouterIsisAddressFamilyIpv6MultiTopologyCfg {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6MultiTopologyCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiTopology = in["multi_topology"].(int)
		ret.Level = in["level"].(string)
		ret.Transition = in["transition"].(int)
		ret.LevelTransition = in["level_transition"].(int)
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6SummaryPrefixList(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6SummaryPrefixList {
	count := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6SummaryPrefixList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6SummaryPrefixList
		oi.Prefix = in["prefix"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6Redistribute(d []interface{}) edpt.RouterIsisAddressFamilyIpv6Redistribute {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6Redistribute
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterIsisAddressFamilyIpv6RedistributeRedistList(in["redist_list"].([]interface{}))
		ret.VipList = getSliceRouterIsisAddressFamilyIpv6RedistributeVipList(in["vip_list"].([]interface{}))
		ret.Isis = getObjectRouterIsisAddressFamilyIpv6RedistributeIsis(in["isis"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6RedistributeRedistList(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList {
	count := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6RedistributeVipList(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6RedistributeVipList {
	count := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6RedistributeVipList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6RedistributeVipList
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		oi.VipMetricType = in["vip_metric_type"].(string)
		oi.VipLevel = in["vip_level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsis(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsis {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsis
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1From = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1From(in["level_1_from"].([]interface{}))
		ret.Level2From = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2From(in["level_2_from"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1From(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Into1 = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1(in["into_1"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1 {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Level2 = in["level_2"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2From(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Into2 = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2(in["into_2"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2 {
	count := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1 = in["level_1"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getObjectRouterIsisAreaPasswordCfg(d []interface{}) edpt.RouterIsisAreaPasswordCfg {
	count := len(d)
	var ret edpt.RouterIsisAreaPasswordCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Password = in["password"].(string)
		ret.Authenticate = getObjectRouterIsisAreaPasswordCfgAuthenticate(in["authenticate"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAreaPasswordCfgAuthenticate(d []interface{}) edpt.RouterIsisAreaPasswordCfgAuthenticate {
	count := len(d)
	var ret edpt.RouterIsisAreaPasswordCfgAuthenticate
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Snp = in["snp"].(string)
	}
	return ret
}

func getObjectRouterIsisAuthentication(d []interface{}) edpt.RouterIsisAuthentication {
	count := len(d)
	var ret edpt.RouterIsisAuthentication
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceRouterIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceRouterIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceRouterIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceRouterIsisAuthenticationSendOnlyList(d []interface{}) []edpt.RouterIsisAuthenticationSendOnlyList {
	count := len(d)
	ret := make([]edpt.RouterIsisAuthenticationSendOnlyList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisAuthenticationModeList(d []interface{}) []edpt.RouterIsisAuthenticationModeList {
	count := len(d)
	ret := make([]edpt.RouterIsisAuthenticationModeList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisAuthenticationKeyChainList(d []interface{}) []edpt.RouterIsisAuthenticationKeyChainList {
	count := len(d)
	ret := make([]edpt.RouterIsisAuthenticationKeyChainList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisDistanceList(d []interface{}) []edpt.RouterIsisDistanceList {
	count := len(d)
	ret := make([]edpt.RouterIsisDistanceList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisDistanceList
		oi.Distance = in["distance"].(int)
		oi.SystemId = in["system_id"].(string)
		oi.Acl = in["acl"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisDomainPasswordCfg(d []interface{}) edpt.RouterIsisDomainPasswordCfg {
	count := len(d)
	var ret edpt.RouterIsisDomainPasswordCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Password = in["password"].(string)
		ret.Authenticate = getObjectRouterIsisDomainPasswordCfgAuthenticate(in["authenticate"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisDomainPasswordCfgAuthenticate(d []interface{}) edpt.RouterIsisDomainPasswordCfgAuthenticate {
	count := len(d)
	var ret edpt.RouterIsisDomainPasswordCfgAuthenticate
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Snp = in["snp"].(string)
	}
	return ret
}

func getSliceRouterIsisHaStandbyExtraCost(d []interface{}) []edpt.RouterIsisHaStandbyExtraCost {
	count := len(d)
	ret := make([]edpt.RouterIsisHaStandbyExtraCost, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisHaStandbyExtraCost
		oi.ExtraCost = in["extra_cost"].(int)
		oi.Group = in["group"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisLogAdjacencyChangesCfg(d []interface{}) edpt.RouterIsisLogAdjacencyChangesCfg {
	count := len(d)
	var ret edpt.RouterIsisLogAdjacencyChangesCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getSliceRouterIsisLspGenIntervalList(d []interface{}) []edpt.RouterIsisLspGenIntervalList {
	count := len(d)
	ret := make([]edpt.RouterIsisLspGenIntervalList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisLspGenIntervalList
		oi.Interval = in["interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisMetricStyleList(d []interface{}) []edpt.RouterIsisMetricStyleList {
	count := len(d)
	ret := make([]edpt.RouterIsisMetricStyleList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisMetricStyleList
		oi.Type = in["type"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisNetList(d []interface{}) []edpt.RouterIsisNetList {
	count := len(d)
	ret := make([]edpt.RouterIsisNetList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisNetList
		oi.Net = in["net"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisPassiveInterfaceList(d []interface{}) []edpt.RouterIsisPassiveInterfaceList {
	count := len(d)
	ret := make([]edpt.RouterIsisPassiveInterfaceList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisPassiveInterfaceList
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Lif = in["lif"].(string)
		oi.Ve = in["ve"].(int)
		oi.Tunnel = in["tunnel"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisProtocolList(d []interface{}) []edpt.RouterIsisProtocolList {
	count := len(d)
	ret := make([]edpt.RouterIsisProtocolList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisProtocolList
		oi.ProtocolTopology = in["protocol_topology"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisRedistribute(d []interface{}) edpt.RouterIsisRedistribute {
	count := len(d)
	var ret edpt.RouterIsisRedistribute
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterIsisRedistributeRedistList(in["redist_list"].([]interface{}))
		ret.VipList = getSliceRouterIsisRedistributeVipList(in["vip_list"].([]interface{}))
		ret.Isis = getObjectRouterIsisRedistributeIsis(in["isis"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIsisRedistributeRedistList(d []interface{}) []edpt.RouterIsisRedistributeRedistList {
	count := len(d)
	ret := make([]edpt.RouterIsisRedistributeRedistList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisRedistributeRedistList
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisRedistributeVipList(d []interface{}) []edpt.RouterIsisRedistributeVipList {
	count := len(d)
	ret := make([]edpt.RouterIsisRedistributeVipList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisRedistributeVipList
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		oi.VipMetricType = in["vip_metric_type"].(string)
		oi.VipLevel = in["vip_level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisRedistributeIsis(d []interface{}) edpt.RouterIsisRedistributeIsis {
	count := len(d)
	var ret edpt.RouterIsisRedistributeIsis
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1From = getObjectRouterIsisRedistributeIsisLevel1From(in["level_1_from"].([]interface{}))
		ret.Level2From = getObjectRouterIsisRedistributeIsisLevel2From(in["level_2_from"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisRedistributeIsisLevel1From(d []interface{}) edpt.RouterIsisRedistributeIsisLevel1From {
	count := len(d)
	var ret edpt.RouterIsisRedistributeIsisLevel1From
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Into1 = getObjectRouterIsisRedistributeIsisLevel1FromInto1(in["into_1"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisRedistributeIsisLevel1FromInto1(d []interface{}) edpt.RouterIsisRedistributeIsisLevel1FromInto1 {
	count := len(d)
	var ret edpt.RouterIsisRedistributeIsisLevel1FromInto1
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Level2 = in["level_2"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getObjectRouterIsisRedistributeIsisLevel2From(d []interface{}) edpt.RouterIsisRedistributeIsisLevel2From {
	count := len(d)
	var ret edpt.RouterIsisRedistributeIsisLevel2From
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Into2 = getObjectRouterIsisRedistributeIsisLevel2FromInto2(in["into_2"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisRedistributeIsisLevel2FromInto2(d []interface{}) edpt.RouterIsisRedistributeIsisLevel2FromInto2 {
	count := len(d)
	var ret edpt.RouterIsisRedistributeIsisLevel2FromInto2
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1 = in["level_1"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getObjectRouterIsisSetOverloadBitCfg(d []interface{}) edpt.RouterIsisSetOverloadBitCfg {
	count := len(d)
	var ret edpt.RouterIsisSetOverloadBitCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SetOverloadBit = in["set_overload_bit"].(int)
		ret.OnStartup = getObjectRouterIsisSetOverloadBitCfgOnStartup(in["on_startup"].([]interface{}))
		ret.SuppressCfg = getObjectRouterIsisSetOverloadBitCfgSuppressCfg(in["suppress_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisSetOverloadBitCfgOnStartup(d []interface{}) edpt.RouterIsisSetOverloadBitCfgOnStartup {
	count := len(d)
	var ret edpt.RouterIsisSetOverloadBitCfgOnStartup
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Delay = in["delay"].(int)
		ret.WaitForBgp = in["wait_for_bgp"].(int)
	}
	return ret
}

func getObjectRouterIsisSetOverloadBitCfgSuppressCfg(d []interface{}) edpt.RouterIsisSetOverloadBitCfgSuppressCfg {
	count := len(d)
	var ret edpt.RouterIsisSetOverloadBitCfgSuppressCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.External = in["external"].(int)
		ret.Interlevel = in["interlevel"].(int)
	}
	return ret
}

func getSliceRouterIsisSpfIntervalExpList(d []interface{}) []edpt.RouterIsisSpfIntervalExpList {
	count := len(d)
	ret := make([]edpt.RouterIsisSpfIntervalExpList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisSpfIntervalExpList
		oi.Min = in["min"].(int)
		oi.Max = in["max"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisSummaryAddressList(d []interface{}) []edpt.RouterIsisSummaryAddressList {
	count := len(d)
	ret := make([]edpt.RouterIsisSummaryAddressList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisSummaryAddressList
		oi.Prefix = in["prefix"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIsis(d *schema.ResourceData) edpt.RouterIsis {
	var ret edpt.RouterIsis
	ret.Inst.AddressFamily = getObjectRouterIsisAddressFamily(d.Get("address_family").([]interface{}))
	ret.Inst.AdjacencyCheck = d.Get("adjacency_check").(int)
	ret.Inst.AreaPasswordCfg = getObjectRouterIsisAreaPasswordCfg(d.Get("area_password_cfg").([]interface{}))
	ret.Inst.Authentication = getObjectRouterIsisAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.Bfd = d.Get("bfd").(string)
	ret.Inst.DefaultInformation = d.Get("default_information").(string)
	ret.Inst.DistanceList = getSliceRouterIsisDistanceList(d.Get("distance_list").([]interface{}))
	ret.Inst.DomainPasswordCfg = getObjectRouterIsisDomainPasswordCfg(d.Get("domain_password_cfg").([]interface{}))
	ret.Inst.HaStandbyExtraCost = getSliceRouterIsisHaStandbyExtraCost(d.Get("ha_standby_extra_cost").([]interface{}))
	ret.Inst.IgnoreLspErrors = d.Get("ignore_lsp_errors").(int)
	ret.Inst.IsType = d.Get("is_type").(string)
	ret.Inst.LogAdjacencyChangesCfg = getObjectRouterIsisLogAdjacencyChangesCfg(d.Get("log_adjacency_changes_cfg").([]interface{}))
	ret.Inst.LspGenIntervalList = getSliceRouterIsisLspGenIntervalList(d.Get("lsp_gen_interval_list").([]interface{}))
	ret.Inst.LspRefreshInterval = d.Get("lsp_refresh_interval").(int)
	ret.Inst.MaxLspLifetime = d.Get("max_lsp_lifetime").(int)
	ret.Inst.MetricStyleList = getSliceRouterIsisMetricStyleList(d.Get("metric_style_list").([]interface{}))
	ret.Inst.NetList = getSliceRouterIsisNetList(d.Get("net_list").([]interface{}))
	ret.Inst.PassiveInterfaceList = getSliceRouterIsisPassiveInterfaceList(d.Get("passive_interface_list").([]interface{}))
	ret.Inst.ProtocolList = getSliceRouterIsisProtocolList(d.Get("protocol_list").([]interface{}))
	ret.Inst.Redistribute = getObjectRouterIsisRedistribute(d.Get("redistribute").([]interface{}))
	ret.Inst.SetOverloadBitCfg = getObjectRouterIsisSetOverloadBitCfg(d.Get("set_overload_bit_cfg").([]interface{}))
	ret.Inst.SpfIntervalExpList = getSliceRouterIsisSpfIntervalExpList(d.Get("spf_interval_exp_list").([]interface{}))
	ret.Inst.SummaryAddressList = getSliceRouterIsisSummaryAddressList(d.Get("summary_address_list").([]interface{}))
	ret.Inst.Tag = d.Get("tag").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
