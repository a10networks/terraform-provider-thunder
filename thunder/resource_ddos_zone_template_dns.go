package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_dns`: DNS template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateDnsCreate,
		UpdateContext: resourceDdosZoneTemplateDnsUpdate,
		ReadContext:   resourceDdosZoneTemplateDnsRead,
		DeleteContext: resourceDdosZoneTemplateDnsDelete,

		Schema: map[string]*schema.Schema{
			"allow_query_class": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_internet_query_class": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "INTERNET query class",
						},
						"allow_csnet_query_class": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "CSNET query class",
						},
						"allow_chaos_query_class": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "CHAOS query class",
						},
						"allow_hesiod_query_class": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HESIOD query class",
						},
						"allow_none_query_class": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NONE query class",
						},
						"allow_any_query_class": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ANY query class",
						},
						"allow_query_class_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take when query class doesn't match",
						},
						"allow_query_class_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"allow_record_type": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_a_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Address record",
						},
						"allow_aaaa_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv6 address record",
						},
						"allow_cname_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Canonical name record",
						},
						"allow_mx_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mail exchange record",
						},
						"allow_ns_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Name server record",
						},
						"allow_srv_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Service locator",
						},
						"record_num_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_num_type": {
										Type: schema.TypeInt, Optional: true, Description: "Other record type value",
									},
								},
							},
						},
						"allow_record_type_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"allow_record_type_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"dns_any_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop DNS queries of Type ANY",
			},
			"dns_any_check_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
			},
			"dns_any_check_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"dns_udp_authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_tcp_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"force_tcp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force DNS request over TCP",
									},
									"force_tcp_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "UDP authentication timeout in seconds",
									},
									"force_tcp_min_delay": {
										Type: schema.TypeInt, Optional: true, Description: "Optional minimum delay (seconds) between DNS retransmits for authentication to pass",
									},
									"force_tcp_ignore_client_source_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow client to retransmit DNS request using different source port during udp-auth (supported in asymmetric mode only)",
									},
								},
							},
						},
						"udp_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "UDP authentication timeout in seconds",
						},
						"min_delay": {
							Type: schema.TypeInt, Optional: true, Description: "Optional minimum delay between DNS retransmits for authentication to pass, unit is specified by min-delay-interval",
						},
						"min_delay_interval": {
							Type: schema.TypeString, Optional: true, Description: "'100ms': 100ms; '1sec': 1sec;",
						},
						"dns_udp_auth_pass_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for passing the authentication",
						},
						"dns_udp_auth_pass_action": {
							Type: schema.TypeString, Optional: true, Description: "'authenticate-src': authenticate-src (Default);",
						},
						"dns_udp_auth_fail_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for failing the authentication. (Applicable to dns-udp retry only)",
						},
						"dns_udp_auth_fail_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src;",
						},
					},
				},
			},
			"domain_group_name": {
				Type: schema.TypeString, Optional: true, Description: "Apply a domain-group to the DNS template",
			},
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fqdn": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_fqdn_rate_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dns_fqdn_rate": {
																Type: schema.TypeInt, Optional: true, Description: "Limiting rate (Range: 5-8000 for FQDN domain based rate limiting, 5-16000000 for FQDN label count based rate limiting)",
															},
															"per": {
																Type: schema.TypeString, Optional: true, Description: "'domain-name': Domain Name; 'src-ip': Source IP address; 'label-count': FQDN label count;",
															},
															"per_domain_per_src_ip": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use both Domain Name and Source IP address for rate-limiting",
															},
															"fqdn_rate_suffix": {
																Type: schema.TypeInt, Optional: true, Description: "Suffix count",
															},
															"fqdn_rate_label_count": {
																Type: schema.TypeInt, Optional: true, Description: "FQDN label count (Range: 1-8)",
															},
														},
													},
												},
												"dns_fqdn_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"dns_fqdn_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
												},
											},
										},
									},
									"domain_group_rate_exceed_action": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop the query (default); 'tunnel-encap-packet': Encapsulate the query and send on a tunnel;",
									},
									"encap_template": {
										Type: schema.TypeString, Optional: true, Description: "DDOS encap template to sepcify the tunnel endpoint",
									},
									"domain_group_rate_per_service": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable per service domain rate checking",
									},
									"request": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"a_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"a": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Address record",
																		},
																		"dns_a_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"aaaa_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aaaa": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv6 address record",
																		},
																		"dns_aaaa_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"cname_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cname": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Canonical name record",
																		},
																		"dns_cname_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"mx_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mx": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mail exchange record",
																		},
																		"dns_mx_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"ns_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ns": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Name server record",
																		},
																		"dns_ns_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"srv_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"srv": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Service locator",
																		},
																		"dns_srv_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "DNS request rate",
																		},
																	},
																},
															},
															"dns_type_cfg": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dns_request_type": {
																			Type: schema.TypeInt, Optional: true, Description: "Other type value",
																		},
																		"dns_request_type_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "request rate limit",
																		},
																	},
																},
															},
														},
													},
												},
												"dst_dns_request_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"dst_dns_request_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
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
			"fqdn_label_count_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of FQDN labels per FQDN",
						},
						"fqdn_label_count_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"fqdn_label_count_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Send reset to client;",
						},
					},
				},
			},
			"fqdn_label_len_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label_length": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum length of FQDN label",
						},
						"fqdn_label_suffix": {
							Type: schema.TypeInt, Optional: true, Description: "Number of suffixes",
						},
						"fqdn_label_length_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"fqdn_label_length_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
					},
				},
			},
			"malformed_query_check": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"validation_type": {
							Type: schema.TypeString, Optional: true, Description: "'basic-header-check': Basic header validation for DNS TCP/UDP queries; 'extended-header-check': Extended header/query validation for DNS TCP/UDP queries; 'disable': Disable Malform query validation for DNS TCP/UDP;",
						},
						"non_query_opcode_check": {
							Type: schema.TypeString, Optional: true, Description: "'disable': When malform check is enabled, TPS always drops DNS query with non query opcode, this option disables this opcode check;",
						},
						"skip_multi_packet_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass DNS fragmented and TCP segmented Queries(Default: dropped)",
						},
						"dns_malformed_query_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"dns_malformed_query_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"multi_pu_threshold_distribution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_pu_threshold_distribution_value": {
							Type: schema.TypeInt, Optional: true, Description: "Destination side rate limit only. Default: 0",
						},
						"multi_pu_threshold_distribution_disable": {
							Type: schema.TypeString, Optional: true, Description: "'disable': Destination side rate limit only. Default: Enable;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"on_no_match": {
				Type: schema.TypeString, Optional: true, Default: "deny", Description: "'permit': permit; 'deny': deny (default);",
			},
			"src": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nxdomain": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_nxdomain_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Limiting rate",
												},
												"dns_nxdomain_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"dns_nxdomain_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
												},
											},
										},
									},
									"request": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"a_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"a": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Address record",
																		},
																		"src_dns_a_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"aaaa_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"aaaa": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv6 address record",
																		},
																		"src_dns_aaaa_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"cname_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"cname": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Canonical name record",
																		},
																		"src_dns_cname_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"mx_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"mx": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mail exchange record",
																		},
																		"src_dns_mx_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"ns_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ns": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Name server record",
																		},
																		"src_dns_ns_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																	},
																},
															},
															"srv_cfg": {
																Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"srv": {
																			Type: schema.TypeInt, Optional: true, Default: 0, Description: "Service locator",
																		},
																		"src_dns_srv_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "DNS request rate",
																		},
																	},
																},
															},
															"dns_type_cfg": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"src_dns_request_type": {
																			Type: schema.TypeInt, Optional: true, Description: "Other type value",
																		},
																		"src_dns_request_type_rate": {
																			Type: schema.TypeInt, Optional: true, Description: "request rate limit",
																		},
																	},
																},
															},
														},
													},
												},
												"src_dns_request_rate_limit_action_list_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
												},
												"src_dns_request_rate_limit_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'reset': Reset client connection;",
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
			"symtimeout_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sym_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Timeout for DNS Symmetric session",
						},
						"sym_timeout_value": {
							Type: schema.TypeInt, Optional: true, Description: "Session timeout value in seconds",
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
func resourceDdosZoneTemplateDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateDnsAllowQueryClass(d []interface{}) edpt.DdosZoneTemplateDnsAllowQueryClass {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsAllowQueryClass
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllowInternetQueryClass = in["allow_internet_query_class"].(int)
		ret.AllowCsnetQueryClass = in["allow_csnet_query_class"].(int)
		ret.AllowChaosQueryClass = in["allow_chaos_query_class"].(int)
		ret.AllowHesiodQueryClass = in["allow_hesiod_query_class"].(int)
		ret.AllowNoneQueryClass = in["allow_none_query_class"].(int)
		ret.AllowAnyQueryClass = in["allow_any_query_class"].(int)
		ret.AllowQueryClassActionListName = in["allow_query_class_action_list_name"].(string)
		ret.AllowQueryClassAction = in["allow_query_class_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsAllowRecordType(d []interface{}) edpt.DdosZoneTemplateDnsAllowRecordType {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsAllowRecordType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllowAType = in["allow_a_type"].(int)
		ret.AllowAaaaType = in["allow_aaaa_type"].(int)
		ret.AllowCnameType = in["allow_cname_type"].(int)
		ret.AllowMxType = in["allow_mx_type"].(int)
		ret.AllowNsType = in["allow_ns_type"].(int)
		ret.AllowSrvType = in["allow_srv_type"].(int)
		ret.RecordNumCfg = getSliceDdosZoneTemplateDnsAllowRecordTypeRecordNumCfg(in["record_num_cfg"].([]interface{}))
		ret.AllowRecordTypeActionListName = in["allow_record_type_action_list_name"].(string)
		ret.AllowRecordTypeAction = in["allow_record_type_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateDnsAllowRecordTypeRecordNumCfg(d []interface{}) []edpt.DdosZoneTemplateDnsAllowRecordTypeRecordNumCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateDnsAllowRecordTypeRecordNumCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateDnsAllowRecordTypeRecordNumCfg
		oi.AllowNumType = in["allow_num_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDnsUdpAuthentication(d []interface{}) edpt.DdosZoneTemplateDnsDnsUdpAuthentication {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDnsUdpAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ForceTcpCfg = getObjectDdosZoneTemplateDnsDnsUdpAuthenticationForceTcpCfg(in["force_tcp_cfg"].([]interface{}))
		ret.UdpTimeout = in["udp_timeout"].(int)
		ret.MinDelay = in["min_delay"].(int)
		ret.MinDelayInterval = in["min_delay_interval"].(string)
		ret.DnsUdpAuthPassActionListName = in["dns_udp_auth_pass_action_list_name"].(string)
		ret.DnsUdpAuthPassAction = in["dns_udp_auth_pass_action"].(string)
		ret.DnsUdpAuthFailActionListName = in["dns_udp_auth_fail_action_list_name"].(string)
		ret.DnsUdpAuthFailAction = in["dns_udp_auth_fail_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDnsUdpAuthenticationForceTcpCfg(d []interface{}) edpt.DdosZoneTemplateDnsDnsUdpAuthenticationForceTcpCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDnsUdpAuthenticationForceTcpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ForceTcp = in["force_tcp"].(int)
		ret.ForceTcpTimeout = in["force_tcp_timeout"].(int)
		ret.ForceTcpMinDelay = in["force_tcp_min_delay"].(int)
		ret.ForceTcpIgnoreClientSourcePort = in["force_tcp_ignore_client_source_port"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDst(d []interface{}) edpt.DdosZoneTemplateDnsDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateDnsDstRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimit(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Fqdn = getObjectDdosZoneTemplateDnsDstRateLimitFqdn(in["fqdn"].([]interface{}))
		ret.DomainGroupRateExceedAction = in["domain_group_rate_exceed_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
		ret.DomainGroupRatePerService = in["domain_group_rate_per_service"].(int)
		ret.Request = getObjectDdosZoneTemplateDnsDstRateLimitRequest(in["request"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitFqdn(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitFqdn {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitFqdn
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsFqdnRateCfg = getSliceDdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg(in["dns_fqdn_rate_cfg"].([]interface{}))
		ret.DnsFqdnRateLimitActionListName = in["dns_fqdn_rate_limit_action_list_name"].(string)
		ret.DnsFqdnRateLimitAction = in["dns_fqdn_rate_limit_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg(d []interface{}) []edpt.DdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg
		oi.DnsFqdnRate = in["dns_fqdn_rate"].(int)
		oi.Per = in["per"].(string)
		oi.PerDomainPerSrcIp = in["per_domain_per_src_ip"].(int)
		oi.FqdnRateSuffix = in["fqdn_rate_suffix"].(int)
		oi.FqdnRateLabelCount = in["fqdn_rate_label_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequest(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequest {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequest
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = getObjectDdosZoneTemplateDnsDstRateLimitRequestType(in["type"].([]interface{}))
		ret.DstDnsRequestRateLimitActionListName = in["dst_dns_request_rate_limit_action_list_name"].(string)
		ret.DstDnsRequestRateLimitAction = in["dst_dns_request_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestType(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestType {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ACfg = getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeACfg(in["a_cfg"].([]interface{}))
		ret.AaaaCfg = getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeAaaaCfg(in["aaaa_cfg"].([]interface{}))
		ret.CnameCfg = getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeCnameCfg(in["cname_cfg"].([]interface{}))
		ret.MxCfg = getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeMxCfg(in["mx_cfg"].([]interface{}))
		ret.NsCfg = getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeNsCfg(in["ns_cfg"].([]interface{}))
		ret.SrvCfg = getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeSrvCfg(in["srv_cfg"].([]interface{}))
		ret.DnsTypeCfg = getSliceDdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg(in["dns_type_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeACfg(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeACfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeACfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.A = in["a"].(int)
		ret.DnsARate = in["dns_a_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeAaaaCfg(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeAaaaCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeAaaaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Aaaa = in["aaaa"].(int)
		ret.DnsAaaaRate = in["dns_aaaa_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeCnameCfg(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeCnameCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeCnameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cname = in["cname"].(int)
		ret.DnsCnameRate = in["dns_cname_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeMxCfg(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeMxCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeMxCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mx = in["mx"].(int)
		ret.DnsMxRate = in["dns_mx_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeNsCfg(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeNsCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeNsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ns = in["ns"].(int)
		ret.DnsNsRate = in["dns_ns_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsDstRateLimitRequestTypeSrvCfg(d []interface{}) edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeSrvCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeSrvCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Srv = in["srv"].(int)
		ret.DnsSrvRate = in["dns_srv_rate"].(int)
	}
	return ret
}

func getSliceDdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg(d []interface{}) []edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg
		oi.DnsRequestType = in["dns_request_type"].(int)
		oi.DnsRequestTypeRate = in["dns_request_type_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsFqdnLabelCountCfg(d []interface{}) edpt.DdosZoneTemplateDnsFqdnLabelCountCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsFqdnLabelCountCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LabelCount = in["label_count"].(int)
		ret.FqdnLabelCountActionListName = in["fqdn_label_count_action_list_name"].(string)
		ret.FqdnLabelCountAction = in["fqdn_label_count_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateDnsFqdnLabelLenCfg(d []interface{}) []edpt.DdosZoneTemplateDnsFqdnLabelLenCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateDnsFqdnLabelLenCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateDnsFqdnLabelLenCfg
		oi.LabelLength = in["label_length"].(int)
		oi.FqdnLabelSuffix = in["fqdn_label_suffix"].(int)
		oi.FqdnLabelLengthActionListName = in["fqdn_label_length_action_list_name"].(string)
		oi.FqdnLabelLengthAction = in["fqdn_label_length_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsMalformedQueryCheck307(d []interface{}) edpt.DdosZoneTemplateDnsMalformedQueryCheck307 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsMalformedQueryCheck307
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ValidationType = in["validation_type"].(string)
		ret.NonQueryOpcodeCheck = in["non_query_opcode_check"].(string)
		ret.SkipMultiPacketCheck = in["skip_multi_packet_check"].(int)
		ret.DnsMalformedQueryActionListName = in["dns_malformed_query_action_list_name"].(string)
		ret.DnsMalformedQueryAction = in["dns_malformed_query_action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateDnsMultiPuThresholdDistribution(d []interface{}) edpt.DdosZoneTemplateDnsMultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsMultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrc(d []interface{}) edpt.DdosZoneTemplateDnsSrc {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimit = getObjectDdosZoneTemplateDnsSrcRateLimit(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimit(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimit {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nxdomain = getObjectDdosZoneTemplateDnsSrcRateLimitNxdomain(in["nxdomain"].([]interface{}))
		ret.Request = getObjectDdosZoneTemplateDnsSrcRateLimitRequest(in["request"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitNxdomain(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitNxdomain {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitNxdomain
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsNxdomainRate = in["dns_nxdomain_rate"].(int)
		ret.DnsNxdomainRateLimitActionListName = in["dns_nxdomain_rate_limit_action_list_name"].(string)
		ret.DnsNxdomainRateLimitAction = in["dns_nxdomain_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequest(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequest {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequest
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = getObjectDdosZoneTemplateDnsSrcRateLimitRequestType(in["type"].([]interface{}))
		ret.SrcDnsRequestRateLimitActionListName = in["src_dns_request_rate_limit_action_list_name"].(string)
		ret.SrcDnsRequestRateLimitAction = in["src_dns_request_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestType(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestType {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ACfg = getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeACfg(in["a_cfg"].([]interface{}))
		ret.AaaaCfg = getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeAaaaCfg(in["aaaa_cfg"].([]interface{}))
		ret.CnameCfg = getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeCnameCfg(in["cname_cfg"].([]interface{}))
		ret.MxCfg = getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeMxCfg(in["mx_cfg"].([]interface{}))
		ret.NsCfg = getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeNsCfg(in["ns_cfg"].([]interface{}))
		ret.SrvCfg = getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeSrvCfg(in["srv_cfg"].([]interface{}))
		ret.DnsTypeCfg = getSliceDdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg(in["dns_type_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeACfg(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeACfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeACfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.A = in["a"].(int)
		ret.SrcDnsARate = in["src_dns_a_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeAaaaCfg(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeAaaaCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeAaaaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Aaaa = in["aaaa"].(int)
		ret.SrcDnsAaaaRate = in["src_dns_aaaa_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeCnameCfg(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeCnameCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeCnameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cname = in["cname"].(int)
		ret.SrcDnsCnameRate = in["src_dns_cname_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeMxCfg(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeMxCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeMxCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mx = in["mx"].(int)
		ret.SrcDnsMxRate = in["src_dns_mx_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeNsCfg(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeNsCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeNsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ns = in["ns"].(int)
		ret.SrcDnsNsRate = in["src_dns_ns_rate"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSrcRateLimitRequestTypeSrvCfg(d []interface{}) edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeSrvCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeSrvCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Srv = in["srv"].(int)
		ret.SrcDnsSrvRate = in["src_dns_srv_rate"].(int)
	}
	return ret
}

func getSliceDdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg(d []interface{}) []edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg
		oi.SrcDnsRequestType = in["src_dns_request_type"].(int)
		oi.SrcDnsRequestTypeRate = in["src_dns_request_type_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateDnsSymtimeoutCfg(d []interface{}) edpt.DdosZoneTemplateDnsSymtimeoutCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateDnsSymtimeoutCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SymTimeout = in["sym_timeout"].(int)
		ret.SymTimeoutValue = in["sym_timeout_value"].(int)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateDns(d *schema.ResourceData) edpt.DdosZoneTemplateDns {
	var ret edpt.DdosZoneTemplateDns
	ret.Inst.AllowQueryClass = getObjectDdosZoneTemplateDnsAllowQueryClass(d.Get("allow_query_class").([]interface{}))
	ret.Inst.AllowRecordType = getObjectDdosZoneTemplateDnsAllowRecordType(d.Get("allow_record_type").([]interface{}))
	ret.Inst.DnsAnyCheck = d.Get("dns_any_check").(int)
	ret.Inst.DnsAnyCheckAction = d.Get("dns_any_check_action").(string)
	ret.Inst.DnsAnyCheckActionListName = d.Get("dns_any_check_action_list_name").(string)
	ret.Inst.DnsUdpAuthentication = getObjectDdosZoneTemplateDnsDnsUdpAuthentication(d.Get("dns_udp_authentication").([]interface{}))
	ret.Inst.DomainGroupName = d.Get("domain_group_name").(string)
	ret.Inst.Dst = getObjectDdosZoneTemplateDnsDst(d.Get("dst").([]interface{}))
	ret.Inst.FqdnLabelCountCfg = getObjectDdosZoneTemplateDnsFqdnLabelCountCfg(d.Get("fqdn_label_count_cfg").([]interface{}))
	ret.Inst.FqdnLabelLenCfg = getSliceDdosZoneTemplateDnsFqdnLabelLenCfg(d.Get("fqdn_label_len_cfg").([]interface{}))
	ret.Inst.MalformedQueryCheck = getObjectDdosZoneTemplateDnsMalformedQueryCheck307(d.Get("malformed_query_check").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosZoneTemplateDnsMultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OnNoMatch = d.Get("on_no_match").(string)
	ret.Inst.Src = getObjectDdosZoneTemplateDnsSrc(d.Get("src").([]interface{}))
	ret.Inst.SymtimeoutCfg = getObjectDdosZoneTemplateDnsSymtimeoutCfg(d.Get("symtimeout_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
