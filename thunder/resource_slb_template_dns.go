package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns`: DNS template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsCreate,
		UpdateContext: resourceSlbTemplateDnsUpdate,
		ReadContext:   resourceSlbTemplateDnsRead,
		DeleteContext: resourceSlbTemplateDnsDelete,

		Schema: map[string]*schema.Schema{
			"add_padding_to_client": {
				Type: schema.TypeString, Optional: true, Description: "'block-length': Block-Length Padding; 'random-block-length': Random-Block-Length Padding;",
			},
			"cache_record_serving_policy": {
				Type: schema.TypeString, Optional: true, Description: "'global': Follow global cofiguration (Default); 'no-change': No change in record order; 'round-robin': Round-robin;",
			},
			"cache_ttl_adjustment_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the ttl adjustment for dns cache response",
			},
			"class_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify a class list name",
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
									"conn_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
									},
									"per": {
										Type: schema.TypeInt, Optional: true, Description: "Per (Number of 100ms)",
									},
									"over_limit_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action when exceeds limit",
									},
									"action_value": {
										Type: schema.TypeString, Optional: true, Description: "'dns-cache-disable': Disable DNS cache when it exceeds limit; 'dns-cache-enable': Enable DNS cache when it exceeds limit; 'forward': Forward the traffic even it exceeds limit;",
									},
									"lockout": {
										Type: schema.TypeInt, Optional: true, Description: "Don't accept any new connection for certain time (Lockout duration in minutes)",
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log a message",
									},
									"log_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Log interval (minute, by default system will log every over limit instance)",
									},
									"dns": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cache_action": {
													Type: schema.TypeString, Optional: true, Default: "cache-disable", Description: "'cache-disable': Disable dns cache; 'cache-enable': Enable dns cache;",
												},
												"ttl": {
													Type: schema.TypeInt, Optional: true, Description: "TTL for cache entry (TTL in seconds)",
												},
												"weight": {
													Type: schema.TypeInt, Optional: true, Description: "Weight for cache entry",
												},
												"honor_server_response_ttl": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Honor the server reponse TTL",
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
			"default_policy": {
				Type: schema.TypeString, Optional: true, Default: "nocache", Description: "'nocache': Cache disable; 'cache': Cache enable;",
			},
			"disable_dns_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS template",
			},
			"disable_ra_cached_resp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS recursive available flag in cached response",
			},
			"disable_rpz_attach_soa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable attaching SOA due to RPZ",
			},
			"dns_logging": {
				Type: schema.TypeString, Optional: true, Description: "dns logging template (DNS Logging template name)",
			},
			"dns64": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS64",
						},
						"cache": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use a cached A-query response to provide AAAA query responses for the same hostname",
						},
						"change_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always change incoming AAAA DNS Query to A",
						},
						"parallel_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward AAAA Query & generate A Query in parallel",
						},
						"retry": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Retry count, default is 3 (Retry Number)",
						},
						"single_response_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Single Response which is used to avoid ambiguity",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Timeout to send additional Queries, unit: second, default is 1",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dnssec_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Use different service group if DNSSEC DO bit set (Service Group Name)",
			},
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop the malformed query",
			},
			"enable_cache_sharing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache sharing",
			},
			"forward": {
				Type: schema.TypeString, Optional: true, Description: "Forward to service group (Service group name)",
			},
			"insert_ipv4": {
				Type: schema.TypeInt, Optional: true, Description: "prefix-length to insert for IPv4",
			},
			"insert_ipv6": {
				Type: schema.TypeInt, Optional: true, Description: "prefix-length to insert for IPv6",
			},
			"local_dns_resolution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostnames": {
										Type: schema.TypeString, Optional: true, Description: "Hostnames class-list name (dns type)",
									},
								},
							},
						},
						"local_resolver_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_resolver": {
										Type: schema.TypeString, Optional: true, Description: "Local dns servers (address)",
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
			"max_cache_entry_size": {
				Type: schema.TypeInt, Optional: true, Default: 1024, Description: "Define maximum cache entry size (Maximum cache entry size per VIP (default 1024))",
			},
			"max_cache_size": {
				Type: schema.TypeInt, Optional: true, Description: "Define maximum cache size (Maximum cache entry per VIP)",
			},
			"max_query_length": {
				Type: schema.TypeInt, Optional: true, Description: "Define Maximum DNS Query Length, default is unlimited (Specify Maximum Length)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Template Name",
			},
			"negative_dns_cache": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_negative_dns_cache": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS negative cache (Need to turn-on the dns-cache for this feature)",
						},
						"bypass_query_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "the threshold bypass the query, default is 100",
						},
						"max_negative_cache_ttl": {
							Type: schema.TypeInt, Optional: true, Default: 7200, Description: "Max negative cache ttl, default is 2 hours",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Period in minutes",
			},
			"query_class_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_class_action": {
							Type: schema.TypeString, Optional: true, Description: "'allow': Allow only certain DNS query classes; 'deny': Deny only certain DNS query classes;",
						},
						"query_class": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_class": {
										Type: schema.TypeString, Optional: true, Description: "'INTERNET': INTERNET query class; 'CHAOS': CHAOS query class; 'HESIOD': HESIOD query class; 'NONE': NONE query class; 'ANY': ANY query class;",
									},
									"num_query_class": {
										Type: schema.TypeInt, Optional: true, Description: "Other query class value",
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
			"query_id_switch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DNS query ID to create sesion",
			},
			"query_type_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_type_action": {
							Type: schema.TypeString, Optional: true, Description: "'allow': Allow only certain DNS query types; 'deny': Deny only certain DNS query types;",
						},
						"query_type": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_type": {
										Type: schema.TypeString, Optional: true, Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;",
									},
									"num_query_type": {
										Type: schema.TypeInt, Optional: true, Description: "Other record type value",
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
			"recursive_dns_resolution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostnames": {
										Type: schema.TypeString, Optional: true, Description: "Hostnames class-list name (dns type), perform resolution while query name matched",
									},
								},
							},
						},
						"csubnet_retry": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "retry when server REFUSED AX inserted EDNS(0) subnet, works only when insert-client-subnet is configured",
						},
						"ns_cache_lookup": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Disable NS Cache Lookup; 'enabled': Enable NS Cache Lookup;",
						},
						"use_service_group_response": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Start Recursive Resolver if Server response doesnt have final answer; 'enabled': Forward Backend Server response to client and dont start recursive resolver;",
						},
						"ipv4_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Source NAT pool or pool group",
						},
						"ipv6_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Source NAT pool or pool group",
						},
						"retries_per_level": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Number of DNS query retries at each server level before closing client connection, default 6",
						},
						"full_response": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Serve all records (authority and additional) when applicable",
						},
						"max_trials": {
							Type: schema.TypeInt, Optional: true, Default: 255, Description: "Total number of times to try DNS query to server before closing client connection, default 255",
						},
						"request_for_pending_resolution": {
							Type: schema.TypeString, Optional: true, Default: "respond-with-servfail", Description: "'drop': Drop of the request during ongoing; 'respond-with-servfail': Respond with SERVFAIL of the request during ongoing; 'start-new-resolution': Start new resolution of the request during ongoing;",
						},
						"udp_retry_interval": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "UDP DNS Retry Interval value 1-6, default is 1 sec (1-6 , default is 1 sec)",
						},
						"udp_initial_interval": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "UDP DNS Retry Interval value 1-6, default is 5 sec (1-6, default is 5sec)",
						},
						"use_client_qid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use client side query id for recursive query",
						},
						"default_recursive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default recursive mode, forward query to bound service-group if hostnames matched",
						},
						"force_cname_resolution": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Force CNAME resolution always; 'disabled': Use answer record in CNAME response if it exists, else resolve;",
						},
						"fast_ns_selection": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Enable fast NS selection; 'disabled': Disable fast NS selection;",
						},
						"dnssec_validation": {
							Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'enabled': Enable DNSSEC validation; 'disabled': Disable DNSSEC validation;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"lookup_order": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query_type": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str_query_type": {
													Type: schema.TypeString, Optional: true, Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;",
												},
												"num_query_type": {
													Type: schema.TypeInt, Optional: true, Description: "Other query type value",
												},
												"order": {
													Type: schema.TypeString, Optional: true, Description: "'ipv4-precede-ipv6': Recursive lookup via IPv4 then IPv6; 'ipv6-precede-ipv4': Recursive lookup via IPv6 then IPv4;",
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
						"gateway_health_check": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query_name": {
										Type: schema.TypeString, Optional: true, Default: "a10networks.com", Description: "Specify the query name used in probe queries, default \"a10networks.com\"",
									},
									"retry": {
										Type: schema.TypeInt, Optional: true, Default: 6, Description: "Maximum number of DNS query retries at each server level before health check fails, default 6 (Retry count (default 6))",
									},
									"timeout": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the health check timeout before retrying or finish, default is 5 sec (Timeout value, in seconds (default 5))",
									},
									"interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the health check interval, default is 10 sec (Interval value, in seconds (default 10))",
									},
									"up_retry": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify number of times that health check consecutively passes before declaring gateway UP, default 1 (up-retry count (default 1))",
									},
									"retry_multi": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify number of times that health check consecutively fails before declaring gateway DOWN, default 1 (retry-multi count (default 1))",
									},
									"gwhc_ns_cache_lookup": {
										Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'disabled': Disable NS Cache Lookup; 'enabled': Enable NS Cache Lookup;",
									},
									"str_query_type": {
										Type: schema.TypeString, Optional: true, Default: "A", Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record;",
									},
									"num_query_type": {
										Type: schema.TypeInt, Optional: true, Description: "Other record type value",
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
			"redirect_to_tcp_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Direct the client to retry with TCP for DNS UDP request",
			},
			"remove_aa_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Make answers created from cache non-authoritative",
			},
			"remove_csubnet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove EDNS(0) client subnet from client queries",
			},
			"remove_padding_to_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove EDNS(0) padding to server",
			},
			"response_rate_limiting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_rate": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
						},
						"filter_response_rate": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Maximum allowed request rate for the filter. This should match average traffic. (default 10 per seconds)",
						},
						"slip_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
						},
						"tc_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will respond with TC bit",
						},
						"match_subnet": {
							Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "IP subnet mask (response rate by IP subnet mask)",
						},
						"match_subnet_v6": {
							Type: schema.TypeInt, Optional: true, Default: 128, Description: "IPV6 subnet mask (response rate by IPv6 subnet mask)",
						},
						"window": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
						},
						"src_ip_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"enable_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"rrl_class_list_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Class-list name",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
									"lid_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lidnum": {
													Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
												},
												"lid_response_rate": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
												},
												"lid_slip_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
												},
												"lid_tc_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will respond with TC bit",
												},
												"lid_match_subnet": {
													Type: schema.TypeString, Optional: true, Default: "255.255.255.255", Description: "IP subnet mask (response rate by IP subnet mask)",
												},
												"lid_match_subnet_v6": {
													Type: schema.TypeInt, Optional: true, Default: 128, Description: "IPV6 subnet mask (response rate by IPv6 subnet mask)",
												},
												"lid_window": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
												},
												"lid_src_ip_only": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
												},
												"lid_enable_log": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
												},
												"lid_action": {
													Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
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
					},
				},
			},
			"rpz_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_id": {
							Type: schema.TypeInt, Required: true, Description: "sequential id of RPZ",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify a Response Policy Zone name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"logging": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log RPZ triggered action",
									},
									"rpz_action": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str_rpz_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Log RPZ due to drop action; 'pass-thru': Log RPZ due to pass-thru action; 'nxdomain': Log RPZ due to nxdomain action; 'nodata': Log RPZ due to nodata action; 'tcp-only': Log RPZ due to tcp-only action; 'local-data': Log RPZ due to local-data action;",
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
			"udp_retransmit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retry_interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "DNS Retry Interval value 1 - 400 in units of 100ms, default is 10 (default is 1000ms) (1 - 400 in units of 100ms, default is 10 (1000ms/1sec))",
						},
						"max_trials": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Total number of times to try DNS query to server before closing client connection, default 3",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceSlbTemplateDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDnsClassList1424(d []interface{}) edpt.SlbTemplateDnsClassList1424 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsClassList1424
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
		ret.LidList = getSliceSlbTemplateDnsClassListLidList1425(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsClassListLidList1425(d []interface{}) []edpt.SlbTemplateDnsClassListLidList1425 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsClassListLidList1425, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsClassListLidList1425
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.Per = in["per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.LogInterval = in["log_interval"].(int)
		oi.Dns = getObjectSlbTemplateDnsClassListLidListDns1426(in["dns"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsClassListLidListDns1426(d []interface{}) edpt.SlbTemplateDnsClassListLidListDns1426 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsClassListLidListDns1426
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func getObjectSlbTemplateDnsDns641427(d []interface{}) edpt.SlbTemplateDnsDns641427 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsDns641427
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Cache = in["cache"].(int)
		ret.ChangeQuery = in["change_query"].(int)
		ret.ParallelQuery = in["parallel_query"].(int)
		ret.Retry = in["retry"].(int)
		ret.SingleResponseDisable = in["single_response_disable"].(int)
		ret.Timeout = in["timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbTemplateDnsLocalDnsResolution1428(d []interface{}) edpt.SlbTemplateDnsLocalDnsResolution1428 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsLocalDnsResolution1428
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostListCfg = getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg1429(in["host_list_cfg"].([]interface{}))
		ret.LocalResolverCfg = getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430(in["local_resolver_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg1429(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg1429 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg1429, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg1429
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430
		oi.LocalResolver = in["local_resolver"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsNegativeDnsCache1431(d []interface{}) edpt.SlbTemplateDnsNegativeDnsCache1431 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsNegativeDnsCache1431
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableNegativeDnsCache = in["enable_negative_dns_cache"].(int)
		ret.BypassQueryThreshold = in["bypass_query_threshold"].(int)
		ret.MaxNegativeCacheTtl = in["max_negative_cache_ttl"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbTemplateDnsQueryClassFilter1432(d []interface{}) edpt.SlbTemplateDnsQueryClassFilter1432 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsQueryClassFilter1432
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryClassAction = in["query_class_action"].(string)
		ret.QueryClass = getSliceSlbTemplateDnsQueryClassFilterQueryClass1433(in["query_class"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsQueryClassFilterQueryClass1433(d []interface{}) []edpt.SlbTemplateDnsQueryClassFilterQueryClass1433 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryClassFilterQueryClass1433, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryClassFilterQueryClass1433
		oi.StrQueryClass = in["str_query_class"].(string)
		oi.NumQueryClass = in["num_query_class"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsQueryTypeFilter1434(d []interface{}) edpt.SlbTemplateDnsQueryTypeFilter1434 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsQueryTypeFilter1434
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryTypeAction = in["query_type_action"].(string)
		ret.QueryType = getSliceSlbTemplateDnsQueryTypeFilterQueryType1435(in["query_type"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsQueryTypeFilterQueryType1435(d []interface{}) []edpt.SlbTemplateDnsQueryTypeFilterQueryType1435 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryTypeFilterQueryType1435, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryTypeFilterQueryType1435
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolution1436(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolution1436 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolution1436
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostListCfg = getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg1437(in["host_list_cfg"].([]interface{}))
		ret.CsubnetRetry = in["csubnet_retry"].(int)
		ret.NsCacheLookup = in["ns_cache_lookup"].(string)
		ret.UseServiceGroupResponse = in["use_service_group_response"].(string)
		ret.Ipv4NatPool = in["ipv4_nat_pool"].(string)
		ret.Ipv6NatPool = in["ipv6_nat_pool"].(string)
		ret.RetriesPerLevel = in["retries_per_level"].(int)
		ret.FullResponse = in["full_response"].(int)
		ret.MaxTrials = in["max_trials"].(int)
		ret.RequestForPendingResolution = in["request_for_pending_resolution"].(string)
		ret.UdpRetryInterval = in["udp_retry_interval"].(int)
		ret.UdpInitialInterval = in["udp_initial_interval"].(int)
		ret.UseClientQid = in["use_client_qid"].(int)
		ret.DefaultRecursive = in["default_recursive"].(int)
		ret.ForceCnameResolution = in["force_cname_resolution"].(string)
		ret.FastNsSelection = in["fast_ns_selection"].(string)
		ret.DnssecValidation = in["dnssec_validation"].(string)
		//omit uuid
		ret.LookupOrder = getObjectSlbTemplateDnsRecursiveDnsResolutionLookupOrder1438(in["lookup_order"].([]interface{}))
		ret.GatewayHealthCheck = getObjectSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1440(in["gateway_health_check"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg1437(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg1437 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg1437, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg1437
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionLookupOrder1438(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder1438 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder1438
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryType = getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439(in["query_type"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		oi.Order = in["order"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1440(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1440 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1440
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryName = in["query_name"].(string)
		ret.Retry = in["retry"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Interval = in["interval"].(int)
		ret.UpRetry = in["up_retry"].(int)
		ret.RetryMulti = in["retry_multi"].(int)
		ret.GwhcNsCacheLookup = in["gwhc_ns_cache_lookup"].(string)
		ret.StrQueryType = in["str_query_type"].(string)
		ret.NumQueryType = in["num_query_type"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbTemplateDnsResponseRateLimiting1441(d []interface{}) edpt.SlbTemplateDnsResponseRateLimiting1441 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsResponseRateLimiting1441
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ResponseRate = in["response_rate"].(int)
		ret.FilterResponseRate = in["filter_response_rate"].(int)
		ret.SlipRate = in["slip_rate"].(int)
		ret.TcRate = in["tc_rate"].(int)
		ret.MatchSubnet = in["match_subnet"].(string)
		ret.MatchSubnetV6 = in["match_subnet_v6"].(int)
		ret.Window = in["window"].(int)
		ret.SrcIpOnly = in["src_ip_only"].(int)
		ret.EnableLog = in["enable_log"].(int)
		ret.Action = in["action"].(string)
		//omit uuid
		ret.RrlClassListList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList1442(in["rrl_class_list_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList1442(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList1442 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList1442, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList1442
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.LidList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443(in["lid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443
		oi.Lidnum = in["lidnum"].(int)
		oi.LidResponseRate = in["lid_response_rate"].(int)
		oi.LidSlipRate = in["lid_slip_rate"].(int)
		oi.LidTcRate = in["lid_tc_rate"].(int)
		oi.LidMatchSubnet = in["lid_match_subnet"].(string)
		oi.LidMatchSubnetV6 = in["lid_match_subnet_v6"].(int)
		oi.LidWindow = in["lid_window"].(int)
		oi.LidSrcIpOnly = in["lid_src_ip_only"].(int)
		oi.LidEnableLog = in["lid_enable_log"].(int)
		oi.LidAction = in["lid_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsRpzList(d []interface{}) []edpt.SlbTemplateDnsRpzList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRpzList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRpzList
		oi.SeqId = in["seq_id"].(int)
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.Logging = getObjectSlbTemplateDnsRpzListLogging(in["logging"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRpzListLogging(d []interface{}) edpt.SlbTemplateDnsRpzListLogging {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRpzListLogging
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.RpzAction = getSliceSlbTemplateDnsRpzListLoggingRpzAction(in["rpz_action"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRpzListLoggingRpzAction(d []interface{}) []edpt.SlbTemplateDnsRpzListLoggingRpzAction {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRpzListLoggingRpzAction, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRpzListLoggingRpzAction
		oi.StrRpzAction = in["str_rpz_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsUdpRetransmit1444(d []interface{}) edpt.SlbTemplateDnsUdpRetransmit1444 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsUdpRetransmit1444
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetryInterval = in["retry_interval"].(int)
		ret.MaxTrials = in["max_trials"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSlbTemplateDns(d *schema.ResourceData) edpt.SlbTemplateDns {
	var ret edpt.SlbTemplateDns
	ret.Inst.AddPaddingToClient = d.Get("add_padding_to_client").(string)
	ret.Inst.CacheRecordServingPolicy = d.Get("cache_record_serving_policy").(string)
	ret.Inst.CacheTtlAdjustmentEnable = d.Get("cache_ttl_adjustment_enable").(int)
	ret.Inst.ClassList = getObjectSlbTemplateDnsClassList1424(d.Get("class_list").([]interface{}))
	ret.Inst.DefaultPolicy = d.Get("default_policy").(string)
	ret.Inst.DisableDnsTemplate = d.Get("disable_dns_template").(int)
	ret.Inst.DisableRaCachedResp = d.Get("disable_ra_cached_resp").(int)
	ret.Inst.DisableRpzAttachSoa = d.Get("disable_rpz_attach_soa").(int)
	ret.Inst.DnsLogging = d.Get("dns_logging").(string)
	ret.Inst.Dns64 = getObjectSlbTemplateDnsDns641427(d.Get("dns64").([]interface{}))
	ret.Inst.DnssecServiceGroup = d.Get("dnssec_service_group").(string)
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.EnableCacheSharing = d.Get("enable_cache_sharing").(int)
	ret.Inst.Forward = d.Get("forward").(string)
	ret.Inst.InsertIpv4 = d.Get("insert_ipv4").(int)
	ret.Inst.InsertIpv6 = d.Get("insert_ipv6").(int)
	ret.Inst.LocalDnsResolution = getObjectSlbTemplateDnsLocalDnsResolution1428(d.Get("local_dns_resolution").([]interface{}))
	ret.Inst.MaxCacheEntrySize = d.Get("max_cache_entry_size").(int)
	ret.Inst.MaxCacheSize = d.Get("max_cache_size").(int)
	ret.Inst.MaxQueryLength = d.Get("max_query_length").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NegativeDnsCache = getObjectSlbTemplateDnsNegativeDnsCache1431(d.Get("negative_dns_cache").([]interface{}))
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.QueryClassFilter = getObjectSlbTemplateDnsQueryClassFilter1432(d.Get("query_class_filter").([]interface{}))
	ret.Inst.QueryIdSwitch = d.Get("query_id_switch").(int)
	ret.Inst.QueryTypeFilter = getObjectSlbTemplateDnsQueryTypeFilter1434(d.Get("query_type_filter").([]interface{}))
	ret.Inst.RecursiveDnsResolution = getObjectSlbTemplateDnsRecursiveDnsResolution1436(d.Get("recursive_dns_resolution").([]interface{}))
	ret.Inst.RedirectToTcpPort = d.Get("redirect_to_tcp_port").(int)
	ret.Inst.RemoveAaFlag = d.Get("remove_aa_flag").(int)
	ret.Inst.RemoveCsubnet = d.Get("remove_csubnet").(int)
	ret.Inst.RemovePaddingToServer = d.Get("remove_padding_to_server").(int)
	ret.Inst.ResponseRateLimiting = getObjectSlbTemplateDnsResponseRateLimiting1441(d.Get("response_rate_limiting").([]interface{}))
	ret.Inst.RpzList = getSliceSlbTemplateDnsRpzList(d.Get("rpz_list").([]interface{}))
	ret.Inst.UdpRetransmit = getObjectSlbTemplateDnsUdpRetransmit1444(d.Get("udp_retransmit").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
