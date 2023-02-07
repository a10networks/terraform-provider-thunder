package thunder

import (
	"context"
	// "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSlbTemplatePolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy`: Policy config\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyCreate,
		UpdateContext: resourceSlbTemplatePolicyUpdate,
		ReadContext:   resourceSlbTemplatePolicyRead,
		DeleteContext: resourceSlbTemplatePolicyDelete,
		Schema: map[string]*schema.Schema{
			"bw_list_id": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Specify id that maps to service group (The id number)",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify a service group (Specify the service group name)",
						},
						"pbslb_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Configure PBSLB logging",
						},
						"pbslb_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3,
							Description: "Specify logging interval in minutes",
						},
						"fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Only log unsuccessful connections",
						},
						"bw_list_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'drop': drop the packet; 'reset': Send reset back;",
						},
						"logging_drp_rst": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Configure PBSLB logging",
						},
						"action_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     3,
							Description: "Specify logging interval in minute (default is 3)",
						},
					},
				},
			},
			"bw_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a blacklist/whitelist name",
			},
			"class_list": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Class list name or geo-location-class-list name",
						},
						"client_ip_l3_dest": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Use destination IP as client IP address",
						},
						"client_ip_l7_header": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Use extract client IP address from L7 header",
						},
						"header_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify L7 header name",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "uuid of the object",
						},
						"lid_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lidnum": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specify a limit ID",
									},
									"conn_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Connection limit",
									},
									"conn_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify connection rate limit",
									},
									"conn_per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Per (Specify interval in number of 100ms)",
									},
									"request_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Request limit (Specify request limit)",
									},
									"request_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Request rate limit (Specify request rate limit)",
									},
									"request_per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Per (Specify interval in number of 100ms)",
									},
									"bw_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify bandwidth rate limit (Bandwidth rate limit in bytes)",
									},
									"bw_per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Per (Specify interval in number of 100ms)",
									},
									"over_limit_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Set action when exceeds limit",
									},
									"action_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;",
									},
									"lockout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Don't accept any new connection for certain time (Lockout duration in minutes)",
									},
									"log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Log a message",
									},
									"interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify log interval in minutes, by default system will log every over limit instance",
									},
									"direct_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Set action when match the lid",
									},
									"direct_service_group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specify a service group (Specify the service group name)",
									},
									"direct_pbslb_logging": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Configure PBSLB logging",
									},
									"direct_pbslb_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     3,
										Description: "Specify logging interval in minutes(default is 3)",
									},
									"direct_fail": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Only log unsuccessful connections",
									},
									"direct_action_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'drop': drop the packet; 'reset': Send reset back;",
									},
									"direct_logging_drp_rst": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Configure PBSLB logging",
									},
									"direct_action_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     3,
										Description: "Specify logging interval in minute (default is 3)",
									},
									"response_code_rate_limit": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code_range_start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "server response code range start",
												},
												"code_range_end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "server response code range end",
												},
												"threshold": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "the times of getting the response code",
												},
												"period": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "seconds",
												},
											},
										},
									},
									"dns64": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Disable",
												},
												"exclusive_answer": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "Exclusive Answer in DNS Response",
												},
												"prefix": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "IPv6 prefix",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"forward_policy": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"no_client_conn_reuse": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Inspects only first request of a connection",
						},
						"acos_event_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Enable acos event logging",
						},
						"local_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Enable local logging",
						},
						"require_web_category": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "Wait for web category to be resolved before taking proxy decision",
						},
						"filtering": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssli_url_filtering": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'bypassed-sni-disable': Disable SNI filtering for bypassed URL's(enabled by default); 'intercepted-sni-enable': Enable SNI filtering for intercepted URL's(disabled by default); 'intercepted-http-disable': Disable HTTP(host/URL) filtering for intercepted URL's(enabled by default); 'no-sni-allow': Allow connection if SNI filtering is enabled and SNI header is not present(Drop by default);",
									},
								},
							},
						},
						"san_filtering": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssli_url_filtering_san": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'enable-san': Enable SAN filtering(disabled by default); 'bypassed-san-disable': Disable SAN filtering for bypassed URL's(enabled by default); 'intercepted-san-enable': Enable SAN filtering for intercepted URL's(disabled by default); 'no-san-allow': Allow connection if SAN filtering is enabled and SAN field is not present(Drop by default);",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "uuid of the object",
						},
						"action_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Action policy name",
									},
									"action1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "'forward-to-internet': Forward request to Internet; 'forward-to-service-group': Forward request to service group; 'forward-to-proxy': Forward request to HTTP proxy server; 'drop': Drop request;",
									},
									"fake_sg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "service group to forward the packets to Internet",
									},
									"real_sg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "service group to forward the packets",
									},
									"forward_snat": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Source NAT pool or pool group",
									},
									"fall_back": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Fallback service group for Internet",
									},
									"fall_back_snat": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Source NAT pool or pool group for fallback server",
									},
									"proxy_chaining": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Enable proxy chaining feature",
									},
									"proxy_chaining_bypass": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Forward all https packets to upstream proxy",
									},
									"support_cert_fetch": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Fetch server certificate by upstream proxy",
									},
									"log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "enable logging",
									},
									"drop_response_code": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Specify response code for drop action",
									},
									"drop_message": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)",
									},
									"drop_redirect_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specify URL to which client request is redirected upon being dropped",
									},
									"http_status_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "302",
										Description: "'301': Moved permanently; '302': Found;",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customized tag",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of requests matching this destination rule;",
												},
											},
										},
									},
								},
							},
						},
						"source_list": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "source destination match rule name",
									},
									"match_class_list": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Class List Name",
									},
									"match_any": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "Match any source",
									},
									"match_authorize_policy": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Authorize-policy for user and group based policy",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Priority of the source(higher the number higher the priority, default 0)",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "uuid of the object",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Customized tag",
									},
									"sampling_enable": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "'all': all; 'hits': Number of requests matching this source rule; 'destination-match-not-found': Number of requests without matching destination rule; 'no-host-info': Failed to parse ip or host information from request;",
												},
											},
										},
									},
									"destination": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list_list": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dest_class_list": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Destination Class List Name",
															},
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Action to be performed",
															},
															"type": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "'host': Match hostname; 'url': Match URL; 'ip': Match destination IP address;",
															},
															"priority": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Priority value of the action(higher the number higher the priority)",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "uuid of the object",
															},
															"sampling_enable": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "'all': all; 'hits': Number of requests matching this destination rule;",
																		},
																	},
																},
															},
														},
													},
												},
												"web_reputation_scope_list": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"web_reputation_scope": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Destination Web Reputation Scope Name",
															},
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Action to be performed",
															},
															"type": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "'host': Match hostname; 'url': match URL;",
															},
															"priority": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Priority value of the action(higher the number higher the priority)",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "uuid of the object",
															},
															"sampling_enable": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "'all': all; 'hits': Number of requests matching this destination rule;",
																		},
																	},
																},
															},
														},
													},
												},
												"web_category_list_list": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"web_category_list": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Destination Web Category List Name",
															},
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Action to be performed",
															},
															"type": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "'host': Match hostname; 'url': match URL;",
															},
															"priority": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "Priority value of the action(higher the number higher the priority)",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "uuid of the object",
															},
															"sampling_enable": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "'all': all; 'hits': Number of requests matching this destination rule;",
																		},
																	},
																},
															},
														},
													},
												},
												"any": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Action to be performed",
															},
															"uuid": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "uuid of the object",
															},
															"sampling_enable": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "'all': all; 'hits': Number of requests matching this destination rule;",
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
			"full_domain_tree": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Share counters between geo-location and sub regions",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Log interval (minute)",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Policy template name",
			},
			"over_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specify operation in case over limit",
			},
			"over_limit_lockup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Don't accept any new connection for certain time (Lockup duration (minute))",
			},
			"over_limit_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Log a message",
			},
			"over_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Reset the connection when it exceeds limit",
			},
			"overlap": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Use overlap mode for geo-location to do longest match",
			},
			"sampling_enable": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'all': all; 'fwd-policy-dns-unresolved': Forward-policy unresolved DNS queries; 'fwd-policy-dns-outstanding': Forward-policy current DNS outstanding requests; 'fwd-policy-snat-fail': Forward-policy source-nat translation failure; 'fwd-policy-hits': Number of forward-policy requests for this policy template; 'fwd-policy-forward-to-internet': Number of forward-policy requests forwarded to internet; 'fwd-policy-forward-to-service-group': Number of forward-policy requests forwarded to service group; 'fwd-policy-forward-to-proxy': Number of forward-policy requests forwarded to proxy; 'fwd-policy-policy-drop': Number of forward-policy requests dropped; 'fwd-policy-source-match-not-found': Forward-policy requests without matching source rule; 'exp-client-hello-not-found': Expected Client HELLO requests not found;",
						},
					},
				},
			},
			"share": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Share counters between virtual ports and virtual servers",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
				Description: "Define timeout value of PBSLB dynamic entry (Timeout value (minute, default is 5))",
			},
			"use_destination_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Use destination IP to match the policy",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customized tag",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "uuid of the object",
			},
		},
	}
}

func resourceSlbTemplatePolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplatePolicyBwListId(d []interface{}) []edpt.SlbTemplatePolicyBwListId {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyBwListId, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyBwListId
		oi.Id = in["id"].(int)
		oi.ServiceGroup = in["service_group"].(string)
		oi.PbslbLogging = in["pbslb_logging"].(int)
		oi.PbslbInterval = in["pbslb_interval"].(int)
		oi.Fail = in["fail"].(int)
		oi.BwListAction = in["bw_list_action"].(string)
		oi.LoggingDrpRst = in["logging_drp_rst"].(int)
		oi.ActionInterval = in["action_interval"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyClassList(d []interface{}) edpt.SlbTemplatePolicyClassList {
	count := len(d)
	var ret edpt.SlbTemplatePolicyClassList
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ClientIpL3Dest = in["client_ip_l3_dest"].(int)
		ret.ClientIpL7Header = in["client_ip_l7_header"].(int)
		ret.HeaderName = in["header_name"].(string)
		//omit uuid
		ret.LidList = getSliceSlbTemplatePolicyClassListLidList(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyClassListLidList(d []interface{}) []edpt.SlbTemplatePolicyClassListLidList {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidList
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnLimit = in["conn_limit"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.ConnPer = in["conn_per"].(int)
		oi.RequestLimit = in["request_limit"].(int)
		oi.RequestRateLimit = in["request_rate_limit"].(int)
		oi.RequestPer = in["request_per"].(int)
		oi.BwRateLimit = in["bw_rate_limit"].(int)
		oi.BwPer = in["bw_per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.Interval = in["interval"].(int)
		oi.DirectAction = in["direct_action"].(int)
		oi.DirectServiceGroup = in["direct_service_group"].(string)
		oi.DirectPbslbLogging = in["direct_pbslb_logging"].(int)
		oi.DirectPbslbInterval = in["direct_pbslb_interval"].(int)
		oi.DirectFail = in["direct_fail"].(int)
		oi.DirectActionValue = in["direct_action_value"].(string)
		oi.DirectLoggingDrpRst = in["direct_logging_drp_rst"].(int)
		oi.DirectActionInterval = in["direct_action_interval"].(int)
		oi.ResponseCodeRateLimit = getSliceSlbTemplatePolicyClassListLidListResponseCodeRateLimit(in["response_code_rate_limit"].([]interface{}))
		oi.Dns64 = getObjectSlbTemplatePolicyClassListLidListDns64(in["dns64"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyClassListLidListResponseCodeRateLimit(d []interface{}) []edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit
		oi.CodeRangeStart = in["code_range_start"].(int)
		oi.CodeRangeEnd = in["code_range_end"].(int)
		oi.Threshold = in["threshold"].(int)
		oi.Period = in["period"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyClassListLidListDns64(d []interface{}) edpt.SlbTemplatePolicyClassListLidListDns64 {
	count := len(d)
	var ret edpt.SlbTemplatePolicyClassListLidListDns64
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicy(d []interface{}) edpt.SlbTemplatePolicyForwardPolicy {
	count := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicy
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.NoClientConnReuse = in["no_client_conn_reuse"].(int)
		ret.AcosEventLog = in["acos_event_log"].(int)
		ret.LocalLogging = in["local_logging"].(int)
		ret.RequireWebCategory = in["require_web_category"].(int)
		ret.Filtering = getSliceSlbTemplatePolicyForwardPolicyFiltering(in["filtering"].([]interface{}))
		ret.SanFiltering = getSliceSlbTemplatePolicyForwardPolicySanFiltering(in["san_filtering"].([]interface{}))
		//omit uuid
		ret.ActionList = getSliceSlbTemplatePolicyForwardPolicyActionList(in["action_list"].([]interface{}))
		ret.SourceList = getSliceSlbTemplatePolicyForwardPolicySourceList(in["source_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyFiltering(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyFiltering {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyFiltering, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyFiltering
		oi.SsliUrlFiltering = in["ssli_url_filtering"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySanFiltering(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySanFiltering {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySanFiltering, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySanFiltering
		oi.SsliUrlFilteringSan = in["ssli_url_filtering_san"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyActionList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionList {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionList
		oi.Name = in["name"].(string)
		oi.Action1 = in["action1"].(string)
		oi.FakeSg = in["fake_sg"].(string)
		oi.RealSg = in["real_sg"].(string)
		oi.ForwardSnat = in["forward_snat"].(string)
		oi.FallBack = in["fall_back"].(string)
		oi.FallBackSnat = in["fall_back_snat"].(string)
		oi.ProxyChaining = in["proxy_chaining"].(int)
		oi.ProxyChainingBypass = in["proxy_chaining_bypass"].(int)
		oi.SupportCertFetch = in["support_cert_fetch"].(int)
		oi.Log = in["log"].(int)
		oi.DropResponseCode = in["drop_response_code"].(int)
		oi.DropMessage = in["drop_message"].(string)
		oi.DropRedirectUrl = in["drop_redirect_url"].(string)
		oi.HttpStatusCode = in["http_status_code"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyActionListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyActionListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceList {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceList
		oi.Name = in["name"].(string)
		oi.MatchClassList = in["match_class_list"].(string)
		oi.MatchAny = in["match_any"].(int)
		oi.MatchAuthorizePolicy = in["match_authorize_policy"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.Destination = getObjectSlbTemplatePolicyForwardPolicySourceListDestination(in["destination"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicySourceListDestination(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceListDestination {
	count := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceListDestination
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListList(in["class_list_list"].([]interface{}))
		ret.WebReputationScopeList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList(in["web_reputation_scope_list"].([]interface{}))
		ret.WebCategoryListList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList(in["web_category_list_list"].([]interface{}))
		ret.Any = getObjectSlbTemplatePolicyForwardPolicySourceListDestinationAny(in["any"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList
		oi.DestClassList = in["dest_class_list"].(string)
		oi.Action = in["action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList
		oi.WebReputationScope = in["web_reputation_scope"].(string)
		oi.Action = in["action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList
		oi.WebCategoryList = in["web_category_list"].(string)
		oi.Action = in["action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicySourceListDestinationAny(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAny {
	count := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAny
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicySamplingEnable(d []interface{}) []edpt.SlbTemplatePolicySamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbTemplatePolicySamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicy(d *schema.ResourceData) edpt.SlbTemplatePolicy {
	var ret edpt.SlbTemplatePolicy
	ret.Inst.BwListId = getSliceSlbTemplatePolicyBwListId(d.Get("bw_list_id").([]interface{}))
	ret.Inst.BwListName = d.Get("bw_list_name").(string)
	ret.Inst.ClassList = getObjectSlbTemplatePolicyClassList(d.Get("class_list").([]interface{}))
	ret.Inst.ForwardPolicy = getObjectSlbTemplatePolicyForwardPolicy(d.Get("forward_policy").([]interface{}))
	ret.Inst.FullDomainTree = d.Get("full_domain_tree").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OverLimit = d.Get("over_limit").(int)
	ret.Inst.OverLimitLockup = d.Get("over_limit_lockup").(int)
	ret.Inst.OverLimitLogging = d.Get("over_limit_logging").(int)
	ret.Inst.OverLimitReset = d.Get("over_limit_reset").(int)
	ret.Inst.Overlap = d.Get("overlap").(int)
	ret.Inst.SamplingEnable = getSliceSlbTemplatePolicySamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Share = d.Get("share").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UseDestinationIp = d.Get("use_destination_ip").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
