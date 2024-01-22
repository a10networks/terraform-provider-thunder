package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id1": {
							Type: schema.TypeInt, Optional: true, Description: "Specify id that maps to service group (The id number)",
						},
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Specify a service group (Specify the service group name)",
						},
						"pbslb_logging": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
						},
						"pbslb_interval": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minutes",
						},
						"fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log unsuccessful connections",
						},
						"bw_list_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': drop the packet; 'reset': Send reset back;",
						},
						"logging_drp_rst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
						},
						"action_interval": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minute (default is 3)",
						},
					},
				},
			},
			"bw_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify a blacklist/whitelist name",
			},
			"class_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Class list name or geo-location-class-list name",
						},
						"client_ip_l3_dest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use destination IP as client IP address",
						},
						"client_ip_l7_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use extract client IP address from L7 header",
						},
						"header_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify L7 header name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"lid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lidnum": {
										Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
									},
									"conn_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Connection limit",
									},
									"conn_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Specify connection rate limit",
									},
									"conn_per": {
										Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
									},
									"request_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Request limit (Specify request limit)",
									},
									"request_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Request rate limit (Specify request rate limit)",
									},
									"request_per": {
										Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
									},
									"bw_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Specify bandwidth rate limit (Bandwidth rate limit in bytes)",
									},
									"bw_per": {
										Type: schema.TypeInt, Optional: true, Description: "Per (Specify interval in number of 100ms)",
									},
									"over_limit_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action when exceeds limit",
									},
									"action_value": {
										Type: schema.TypeString, Optional: true, Description: "'forward': Forward the traffic even it exceeds limit; 'reset': Reset the connection when it exceeds limit;",
									},
									"lockout": {
										Type: schema.TypeInt, Optional: true, Description: "Don't accept any new connection for certain time (Lockout duration in minutes)",
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log a message",
									},
									"interval": {
										Type: schema.TypeInt, Optional: true, Description: "Specify log interval in minutes, by default system will log every over limit instance",
									},
									"direct_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set action when match the lid",
									},
									"direct_service_group": {
										Type: schema.TypeString, Optional: true, Description: "Specify a service group (Specify the service group name)",
									},
									"direct_pbslb_logging": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
									},
									"direct_pbslb_interval": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minutes(default is 3)",
									},
									"direct_fail": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log unsuccessful connections",
									},
									"direct_action_value": {
										Type: schema.TypeString, Optional: true, Description: "'drop': drop the packet; 'reset': Send reset back;",
									},
									"direct_logging_drp_rst": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure PBSLB logging",
									},
									"direct_action_interval": {
										Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify logging interval in minute (default is 3)",
									},
									"response_code_rate_limit": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code_range_start": {
													Type: schema.TypeInt, Optional: true, Description: "server response code range start",
												},
												"code_range_end": {
													Type: schema.TypeInt, Optional: true, Description: "server response code range end",
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "the times of getting the response code",
												},
												"period": {
													Type: schema.TypeInt, Optional: true, Description: "seconds",
												},
											},
										},
									},
									"dns64": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable",
												},
												"exclusive_answer": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Exclusive Answer in DNS Response",
												},
												"prefix": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
												},
											},
										},
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
					},
				},
			},
			"forward_policy": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"no_client_conn_reuse": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inspects only first request of a connection",
						},
						"acos_event_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable acos event logging",
						},
						"local_logging": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable local logging",
						},
						"require_web_category": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wait for web category to be resolved before taking proxy decision",
						},
						"forward_http_connect_to_icap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward HTTP CONNECT request to ICAP server",
						},
						"reqmod_icap": {
							Type: schema.TypeString, Optional: true, Description: "ICAP reqmod template (Reqmod ICAP Template Name)",
						},
						"filtering": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssli_url_filtering": {
										Type: schema.TypeString, Optional: true, Description: "'bypassed-sni-disable': Disable SNI filtering for bypassed URL's(enabled by default); 'intercepted-sni-enable': Enable SNI filtering for intercepted URL's(disabled by default); 'intercepted-http-disable': Disable HTTP(host/URL) filtering for intercepted URL's(enabled by default); 'no-sni-allow': Allow connection if SNI filtering is enabled and SNI header is not present(Drop by default);",
									},
								},
							},
						},
						"san_filtering": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssli_url_filtering_san": {
										Type: schema.TypeString, Optional: true, Description: "'enable-san': Enable SAN filtering(disabled by default); 'bypassed-san-disable': Disable SAN filtering for bypassed URL's(enabled by default); 'intercepted-san-enable': Enable SAN filtering for intercepted URL's(disabled by default); 'no-san-allow': Allow connection if SAN filtering is enabled and SAN field is not present(Drop by default);",
									},
								},
							},
						},
						"enable_adv_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable adv-match rules and deactive all the other kinds of destination rules",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"action_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Action policy name",
									},
									"action1": {
										Type: schema.TypeString, Optional: true, Description: "'forward-to-internet': Forward request to Internet; 'forward-to-service-group': Forward request to service group; 'forward-to-proxy': Forward request to HTTP proxy server; 'drop': Drop request;",
									},
									"fake_sg": {
										Type: schema.TypeString, Optional: true, Description: "service group to forward the packets to Internet",
									},
									"real_sg": {
										Type: schema.TypeString, Optional: true, Description: "service group to forward the packets",
									},
									"forward_snat": {
										Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group",
									},
									"forward_snat_pt_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source port translation only",
									},
									"fall_back": {
										Type: schema.TypeString, Optional: true, Description: "Fallback service group for Internet",
									},
									"fall_back_snat": {
										Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group for fallback server",
									},
									"fall_back_snat_pt_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source port translation only for fallback server",
									},
									"proxy_chaining": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable proxy chaining feature",
									},
									"proxy_chaining_bypass": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward all https packets to upstream proxy",
									},
									"support_cert_fetch": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fetch server certificate by upstream proxy",
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable logging",
									},
									"drop_response_code": {
										Type: schema.TypeInt, Optional: true, Description: "Specify response code for drop action",
									},
									"drop_message": {
										Type: schema.TypeString, Optional: true, Description: "drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)",
									},
									"drop_redirect_url": {
										Type: schema.TypeString, Optional: true, Description: "Specify URL to which client request is redirected upon being dropped",
									},
									"http_status_code": {
										Type: schema.TypeString, Optional: true, Default: "302", Description: "'301': Moved permanently; '302': Found;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this destination rule;",
												},
											},
										},
									},
								},
							},
						},
						"dual_stack_action_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Action name",
									},
									"ipv4": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 service group to forward",
									},
									"ipv4_snat": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 source NAT pool or pool group",
									},
									"ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 service group to forward",
									},
									"ipv6_snat": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 source NAT pool or pool group",
									},
									"fall_back": {
										Type: schema.TypeString, Optional: true, Description: "Fallback service group",
									},
									"fall_back_snat": {
										Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group for fallback",
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable logging",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests forward by this action;",
												},
											},
										},
									},
								},
							},
						},
						"source_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "source destination match rule name",
									},
									"match_class_list": {
										Type: schema.TypeString, Optional: true, Description: "Class List Name",
									},
									"match_any": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match any source",
									},
									"match_authorize_policy": {
										Type: schema.TypeString, Optional: true, Description: "Authorize-policy for user and group based policy",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Description: "Priority of the source(higher the number higher the priority, default 0)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"sampling_enable": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this source rule; 'destination-match-not-found': Number of requests without matching destination rule; 'no-host-info': Failed to parse ip or host information from request;",
												},
											},
										},
									},
									"destination": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"adv_match_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"priority": {
																Type: schema.TypeInt, Required: true, Description: "Rule priority (1000 is highest)",
															},
															"match_host": {
																Type: schema.TypeString, Optional: true, Description: "Match request host (HTTP stage) or SNI/SAN (SSL stage)",
															},
															"match_http_content_encoding": {
																Type: schema.TypeString, Optional: true, Description: "Match the value of HTTP header \"Content-Encoding\"",
															},
															"match_http_content_length_range_begin": {
																Type: schema.TypeInt, Optional: true, Description: "Match the value of HTTP header \"Content-Length\" with an inclusive range",
															},
															"match_http_content_length_range_end": {
																Type: schema.TypeInt, Optional: true, Description: "End of the \"Content-Length\" range",
															},
															"match_http_content_type": {
																Type: schema.TypeString, Optional: true, Description: "Match the value of HTTP header \"Content-Type\"",
															},
															"match_http_header": {
																Type: schema.TypeString, Optional: true, Description: "Matching the name of all request headers",
															},
															"match_http_method_connect": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method CONNECT",
															},
															"match_http_method_delete": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method DELETE",
															},
															"match_http_method_get": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method GET",
															},
															"match_http_method_head": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method HEAD",
															},
															"match_http_method_options": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method OPTIONS",
															},
															"match_http_method_patch": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method PATCH",
															},
															"match_http_method_post": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method POST",
															},
															"match_http_method_put": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method PUT",
															},
															"match_http_method_trace": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Match HTTP request method TRACE",
															},
															"match_http_request_file_extension": {
																Type: schema.TypeString, Optional: true, Description: "Match file extension of URL in HTTP request line",
															},
															"match_http_url_regex": {
																Type: schema.TypeString, Optional: true, Description: "Match URI in HTTP request line by given regular expression",
															},
															"match_http_url": {
																Type: schema.TypeString, Optional: true, Description: "Match URL in HTTP request line",
															},
															"match_http_user_agent": {
																Type: schema.TypeString, Optional: true, Description: "Matching the value of HTTP header \"User-Agent\"",
															},
															"match_server_address": {
																Type: schema.TypeString, Optional: true, Description: "Match target server IP address",
															},
															"match_server_port": {
																Type: schema.TypeInt, Optional: true, Description: "Match target server port number",
															},
															"match_server_port_range_begin": {
																Type: schema.TypeInt, Optional: true, Description: "Math targer server port range inclusively",
															},
															"match_server_port_range_end": {
																Type: schema.TypeInt, Optional: true, Description: "End of port range",
															},
															"match_time_range": {
																Type: schema.TypeString, Optional: true, Description: "Enable rule in this time-range",
															},
															"match_web_category_list": {
																Type: schema.TypeString, Optional: true, Description: "Match web-category list",
															},
															"match_web_reputation_scope": {
																Type: schema.TypeString, Optional: true, Description: "Match web-reputation scope",
															},
															"disable_reqmod_icap": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable REQMOD ICAP template",
															},
															"disable_respmod_icap": {
																Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable RESPMOD ICAP template",
															},
															"notify_page": {
																Type: schema.TypeString, Optional: true, Description: "Send notify-page to client",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "Forwading action of this rule",
															},
															"dual_stack_action": {
																Type: schema.TypeString, Optional: true, Description: "Forwarding action of this rule",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"user_tag": {
																Type: schema.TypeString, Optional: true, Description: "Customized tag",
															},
															"sampling_enable": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests hit this rule;",
																		},
																	},
																},
															},
														},
													},
												},
												"class_list_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dest_class_list": {
																Type: schema.TypeString, Required: true, Description: "Destination Class List Name",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "Action to be performed",
															},
															"dual_stack_action": {
																Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
															},
															"type": {
																Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': Match URL; 'ip': Match destination IP address;",
															},
															"priority": {
																Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"web_reputation_scope_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"web_reputation_scope": {
																Type: schema.TypeString, Required: true, Description: "Destination Web Reputation Scope Name",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "Action to be performed",
															},
															"dual_stack_action": {
																Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
															},
															"type": {
																Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': match URL;",
															},
															"priority": {
																Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"web_category_list_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"web_category_list": {
																Type: schema.TypeString, Required: true, Description: "Destination Web Category List Name",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "Action to be performed",
															},
															"dual_stack_action": {
																Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
															},
															"type": {
																Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': match URL;",
															},
															"priority": {
																Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"any": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type: schema.TypeString, Optional: true, Description: "Action to be performed",
															},
															"dual_stack_action": {
																Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
															"sampling_enable": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of requests matching this destination rule;",
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share counters between geo-location and sub regions",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Log interval (minute)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Policy template name",
			},
			"over_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify operation in case over limit",
			},
			"over_limit_lockup": {
				Type: schema.TypeInt, Optional: true, Description: "Don't accept any new connection for certain time (Lockup duration (minute))",
			},
			"over_limit_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log a message",
			},
			"over_limit_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset the connection when it exceeds limit",
			},
			"overlap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use overlap mode for geo-location to do longest match",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'fwd-policy-dns-unresolved': Forward-policy unresolved DNS queries; 'fwd-policy-dns-outstanding': Forward-policy current DNS outstanding requests; 'fwd-policy-snat-fail': Forward-policy source-nat translation failure; 'fwd-policy-hits': Number of forward-policy requests for this policy template; 'fwd-policy-forward-to-internet': Number of forward-policy requests forwarded to internet; 'fwd-policy-forward-to-service-group': Number of forward-policy requests forwarded to service group; 'fwd-policy-forward-to-proxy': Number of forward-policy requests forwarded to proxy; 'fwd-policy-policy-drop': Number of forward-policy requests dropped; 'fwd-policy-source-match-not-found': Forward-policy requests without matching source rule; 'exp-client-hello-not-found': Expected Client HELLO requests not found;",
						},
					},
				},
			},
			"share": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share counters between virtual ports and virtual servers",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Define timeout value of PBSLB dynamic entry (Timeout value (minute, default is 5))",
			},
			"use_destination_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use destination IP to match the policy",
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

func getSliceSlbTemplatePolicyBwListId(d []interface{}) []edpt.SlbTemplatePolicyBwListId {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyBwListId, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyBwListId
		oi.Id1 = in["id1"].(int)
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

func getObjectSlbTemplatePolicyClassList1453(d []interface{}) edpt.SlbTemplatePolicyClassList1453 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyClassList1453
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ClientIpL3Dest = in["client_ip_l3_dest"].(int)
		ret.ClientIpL7Header = in["client_ip_l7_header"].(int)
		ret.HeaderName = in["header_name"].(string)
		//omit uuid
		ret.LidList = getSliceSlbTemplatePolicyClassListLidList1454(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyClassListLidList1454(d []interface{}) []edpt.SlbTemplatePolicyClassListLidList1454 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidList1454, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidList1454
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
		oi.ResponseCodeRateLimit = getSliceSlbTemplatePolicyClassListLidListResponseCodeRateLimit1455(in["response_code_rate_limit"].([]interface{}))
		oi.Dns64 = getObjectSlbTemplatePolicyClassListLidListDns641456(in["dns64"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyClassListLidListResponseCodeRateLimit1455(d []interface{}) []edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit1455 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit1455, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyClassListLidListResponseCodeRateLimit1455
		oi.CodeRangeStart = in["code_range_start"].(int)
		oi.CodeRangeEnd = in["code_range_end"].(int)
		oi.Threshold = in["threshold"].(int)
		oi.Period = in["period"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyClassListLidListDns641456(d []interface{}) edpt.SlbTemplatePolicyClassListLidListDns641456 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyClassListLidListDns641456
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Disable = in["disable"].(int)
		ret.ExclusiveAnswer = in["exclusive_answer"].(int)
		ret.Prefix = in["prefix"].(string)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicy1457(d []interface{}) edpt.SlbTemplatePolicyForwardPolicy1457 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicy1457
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NoClientConnReuse = in["no_client_conn_reuse"].(int)
		ret.AcosEventLog = in["acos_event_log"].(int)
		ret.LocalLogging = in["local_logging"].(int)
		ret.RequireWebCategory = in["require_web_category"].(int)
		ret.ForwardHttpConnectToIcap = in["forward_http_connect_to_icap"].(int)
		ret.ReqmodIcap = in["reqmod_icap"].(string)
		ret.Filtering = getSliceSlbTemplatePolicyForwardPolicyFiltering1458(in["filtering"].([]interface{}))
		ret.SanFiltering = getSliceSlbTemplatePolicyForwardPolicySanFiltering1459(in["san_filtering"].([]interface{}))
		ret.EnableAdvMatch = in["enable_adv_match"].(int)
		//omit uuid
		ret.ActionList = getSliceSlbTemplatePolicyForwardPolicyActionList1460(in["action_list"].([]interface{}))
		ret.DualStackActionList = getSliceSlbTemplatePolicyForwardPolicyDualStackActionList1462(in["dual_stack_action_list"].([]interface{}))
		ret.SourceList = getSliceSlbTemplatePolicyForwardPolicySourceList1464(in["source_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyFiltering1458(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyFiltering1458 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyFiltering1458, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyFiltering1458
		oi.SsliUrlFiltering = in["ssli_url_filtering"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySanFiltering1459(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySanFiltering1459 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySanFiltering1459, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySanFiltering1459
		oi.SsliUrlFilteringSan = in["ssli_url_filtering_san"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyActionList1460(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionList1460 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionList1460, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionList1460
		oi.Name = in["name"].(string)
		oi.Action1 = in["action1"].(string)
		oi.FakeSg = in["fake_sg"].(string)
		oi.RealSg = in["real_sg"].(string)
		oi.ForwardSnat = in["forward_snat"].(string)
		oi.ForwardSnatPtOnly = in["forward_snat_pt_only"].(int)
		oi.FallBack = in["fall_back"].(string)
		oi.FallBackSnat = in["fall_back_snat"].(string)
		oi.FallBackSnatPtOnly = in["fall_back_snat_pt_only"].(int)
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
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyActionListSamplingEnable1461(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyActionListSamplingEnable1461(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable1461 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable1461, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyActionListSamplingEnable1461
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyDualStackActionList1462(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyDualStackActionList1462 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyDualStackActionList1462, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyDualStackActionList1462
		oi.Name = in["name"].(string)
		oi.Ipv4 = in["ipv4"].(string)
		oi.Ipv4Snat = in["ipv4_snat"].(string)
		oi.Ipv6 = in["ipv6"].(string)
		oi.Ipv6Snat = in["ipv6_snat"].(string)
		oi.FallBack = in["fall_back"].(string)
		oi.FallBackSnat = in["fall_back_snat"].(string)
		oi.Log = in["log"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceList1464(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceList1464 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceList1464, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceList1464
		oi.Name = in["name"].(string)
		oi.MatchClassList = in["match_class_list"].(string)
		oi.MatchAny = in["match_any"].(int)
		oi.MatchAuthorizePolicy = in["match_authorize_policy"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListSamplingEnable1465(in["sampling_enable"].([]interface{}))
		oi.Destination = getObjectSlbTemplatePolicyForwardPolicySourceListDestination1466(in["destination"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListSamplingEnable1465(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable1465 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable1465, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListSamplingEnable1465
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicySourceListDestination1466(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceListDestination1466 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceListDestination1466
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdvMatchList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467(in["adv_match_list"].([]interface{}))
		ret.ClassListList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469(in["class_list_list"].([]interface{}))
		ret.WebReputationScopeList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470(in["web_reputation_scope_list"].([]interface{}))
		ret.WebCategoryListList = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471(in["web_category_list_list"].([]interface{}))
		ret.Any = getObjectSlbTemplatePolicyForwardPolicySourceListDestinationAny1472(in["any"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467
		oi.Priority = in["priority"].(int)
		oi.MatchHost = in["match_host"].(string)
		oi.MatchHttpContentEncoding = in["match_http_content_encoding"].(string)
		oi.MatchHttpContentLengthRangeBegin = in["match_http_content_length_range_begin"].(int)
		oi.MatchHttpContentLengthRangeEnd = in["match_http_content_length_range_end"].(int)
		oi.MatchHttpContentType = in["match_http_content_type"].(string)
		oi.MatchHttpHeader = in["match_http_header"].(string)
		oi.MatchHttpMethodConnect = in["match_http_method_connect"].(int)
		oi.MatchHttpMethodDelete = in["match_http_method_delete"].(int)
		oi.MatchHttpMethodGet = in["match_http_method_get"].(int)
		oi.MatchHttpMethodHead = in["match_http_method_head"].(int)
		oi.MatchHttpMethodOptions = in["match_http_method_options"].(int)
		oi.MatchHttpMethodPatch = in["match_http_method_patch"].(int)
		oi.MatchHttpMethodPost = in["match_http_method_post"].(int)
		oi.MatchHttpMethodPut = in["match_http_method_put"].(int)
		oi.MatchHttpMethodTrace = in["match_http_method_trace"].(int)
		oi.MatchHttpRequestFileExtension = in["match_http_request_file_extension"].(string)
		oi.MatchHttpUrlRegex = in["match_http_url_regex"].(string)
		oi.MatchHttpUrl = in["match_http_url"].(string)
		oi.MatchHttpUserAgent = in["match_http_user_agent"].(string)
		oi.MatchServerAddress = in["match_server_address"].(string)
		oi.MatchServerPort = in["match_server_port"].(int)
		oi.MatchServerPortRangeBegin = in["match_server_port_range_begin"].(int)
		oi.MatchServerPortRangeEnd = in["match_server_port_range_end"].(int)
		oi.MatchTimeRange = in["match_time_range"].(string)
		oi.MatchWebCategoryList = in["match_web_category_list"].(string)
		oi.MatchWebReputationScope = in["match_web_reputation_scope"].(string)
		oi.DisableReqmodIcap = in["disable_reqmod_icap"].(int)
		oi.DisableRespmodIcap = in["disable_respmod_icap"].(int)
		oi.NotifyPage = in["notify_page"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469
		oi.DestClassList = in["dest_class_list"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470
		oi.WebReputationScope = in["web_reputation_scope"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471
		oi.WebCategoryList = in["web_category_list"].(string)
		oi.Action = in["action"].(string)
		oi.DualStackAction = in["dual_stack_action"].(string)
		oi.Type = in["type"].(string)
		oi.Priority = in["priority"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplatePolicyForwardPolicySourceListDestinationAny1472(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAny1472 {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAny1472
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.DualStackAction = in["dual_stack_action"].(string)
		//omit uuid
		ret.SamplingEnable = getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473(d []interface{}) []edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplatePolicySamplingEnable(d []interface{}) []edpt.SlbTemplatePolicySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplatePolicySamplingEnable, 0, count1)
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
	ret.Inst.ClassList = getObjectSlbTemplatePolicyClassList1453(d.Get("class_list").([]interface{}))
	ret.Inst.ForwardPolicy = getObjectSlbTemplatePolicyForwardPolicy1457(d.Get("forward_policy").([]interface{}))
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
