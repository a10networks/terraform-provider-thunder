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
			"cache_hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache entry hit count",
			},
			"cache_record_serving_policy": {
				Type: schema.TypeString, Optional: true, Description: "'global': Follow global cofiguration (Default); 'no-change': No change in record order; 'round-robin': Round-robin;",
			},
			"cache_ttl_adjustment_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the ttl adjustment for dns cache response",
			},
			"category_lookup_bypass": {
				Type: schema.TypeString, Optional: true, Description: "DNS type class-lists for bypassing category lookup",
			},
			"category_lookup_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category_name": {
							Type: schema.TypeString, Required: true, Description: "category-list name",
						},
						"permit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Permit matching DNS domains",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Deny matching DNS domains",
						},
						"respond": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond to matching DNS domains",
						},
						"respond_nxdomain": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond with NXDOMAIN",
						},
						"respond_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Type A record to respond (IPv4 address)",
						},
						"respond_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "TYPE AAAA record to respond (IPv6 address)",
						},
						"respond_cname_str": {
							Type: schema.TypeString, Optional: true, Description: "CNAME to respond (Canonical name)",
						},
						"response_ttl": {
							Type: schema.TypeInt, Optional: true, Default: 300, Description: "Set response TTL in seconds (TTL value in seconds)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"category_lookup_online_lookup": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable online webroot lookup",
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
			"dns_cookie_cache_policy": {
				Type: schema.TypeString, Optional: true, Description: "'served-by-cache': Answer from cache for requests with cookie; 'served-by-backend': Answer from server for requests with cookie;",
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
			"label_count_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop_log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the log when hit the rule",
						},
						"label_count_filter_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': drop; 'ignore': ignore;",
						},
						"min_fqdn_label_count": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum number of FQDN labels per FQDN",
						},
						"max_fqdn_label_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of FQDN labels per FQDN",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"label_length_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop_log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the log when hit the rule",
						},
						"label_length_filter_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': drop; 'ignore': ignore;",
						},
						"fqdn_label_length": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"length": {
										Type: schema.TypeInt, Optional: true, Description: "fqdn label length",
									},
									"suffix": {
										Type: schema.TypeInt, Optional: true, Description: "",
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
			"max_udp_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set maximum DNS response message size that ACOS sends by UDP (Maximum DNS response message size (bytes))",
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
						"cache_non_valid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable caching non-valid negative response, otherwise will only cache valid negative response",
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
			"qps_log_high": {
				Type: schema.TypeInt, Optional: true, Description: "high threshold for QPS logging (queries per second)",
			},
			"qps_log_low": {
				Type: schema.TypeInt, Optional: true, Description: "low threshold for QPS logging (queries per second)",
			},
			"qps_threshold_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable threshold log for DNS QPS",
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
						"ns_longest_match": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Look up NS of top level label, do a nearly-full resolution; 'enabled': Enable NS cache longest match;",
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
						"parallel_queries": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Number of parallel queries to send to servers",
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
						"edns_udp_size": {
							Type: schema.TypeInt, Optional: true, Default: 4096, Description: "Set EDNS UDP payload size of queries sent during resolution (EDNS UDP payload size of queries, default:4096 bytes)",
						},
						"max_signature_validation_attempts": {
							Type: schema.TypeInt, Optional: true, Description: "Set maximum number of times DNSSEC signature validation attempts allowed per resolution",
						},
						"max_signature_validation_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Set maximum number of times DNSSEC signature validation failures allowed per resolution",
						},
						"max_key_digest_validation_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Set maximum number of times DNSSEC key-digest validation failures allowed per resolution",
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
						"nx_response_rate": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Queries from entries whose NX Responses exceeding this rate within the window will be dropped (default 5 per second)",
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
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second), 0 for unlimited",
												},
												"lid_slip_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
												},
												"lid_nx_response_rate": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Queries from entries whose NX Responses exceeding this rate within the window will be dropped (default 5 per second)",
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
			"tld_filter_log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable dns tld filter logging",
			},
			"tld_filter_white_list": {
				Type: schema.TypeString, Optional: true, Description: "white-list class-list name (string-insensitive type)",
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

func getSliceSlbTemplateDnsCategoryLookupList(d []interface{}) []edpt.SlbTemplateDnsCategoryLookupList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsCategoryLookupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsCategoryLookupList
		oi.CategoryName = in["category_name"].(string)
		oi.Permit = in["permit"].(int)
		oi.Drop = in["drop"].(int)
		oi.Respond = in["respond"].(int)
		oi.RespondNxdomain = in["respond_nxdomain"].(int)
		oi.RespondIpAddr = in["respond_ip_addr"].(string)
		oi.RespondIpv6Addr = in["respond_ipv6_addr"].(string)
		oi.RespondCnameStr = in["respond_cname_str"].(string)
		oi.ResponseTtl = in["response_ttl"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsClassList1524(d []interface{}) edpt.SlbTemplateDnsClassList1524 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsClassList1524
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
		ret.LidList = getSliceSlbTemplateDnsClassListLidList1525(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsClassListLidList1525(d []interface{}) []edpt.SlbTemplateDnsClassListLidList1525 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsClassListLidList1525, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsClassListLidList1525
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.Per = in["per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.LogInterval = in["log_interval"].(int)
		oi.Dns = getObjectSlbTemplateDnsClassListLidListDns1526(in["dns"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsClassListLidListDns1526(d []interface{}) edpt.SlbTemplateDnsClassListLidListDns1526 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsClassListLidListDns1526
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func getObjectSlbTemplateDnsDns641527(d []interface{}) edpt.SlbTemplateDnsDns641527 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsDns641527
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

func getObjectSlbTemplateDnsLabelCountFilter1528(d []interface{}) edpt.SlbTemplateDnsLabelCountFilter1528 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsLabelCountFilter1528
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DropLogEnable = in["drop_log_enable"].(int)
		ret.LabelCountFilterAction = in["label_count_filter_action"].(string)
		ret.MinFqdnLabelCount = in["min_fqdn_label_count"].(int)
		ret.MaxFqdnLabelCount = in["max_fqdn_label_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbTemplateDnsLabelLengthFilter1529(d []interface{}) edpt.SlbTemplateDnsLabelLengthFilter1529 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsLabelLengthFilter1529
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DropLogEnable = in["drop_log_enable"].(int)
		ret.LabelLengthFilterAction = in["label_length_filter_action"].(string)
		ret.FqdnLabelLength = getSliceSlbTemplateDnsLabelLengthFilterFqdnLabelLength1530(in["fqdn_label_length"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsLabelLengthFilterFqdnLabelLength1530(d []interface{}) []edpt.SlbTemplateDnsLabelLengthFilterFqdnLabelLength1530 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLabelLengthFilterFqdnLabelLength1530, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLabelLengthFilterFqdnLabelLength1530
		oi.Length = in["length"].(int)
		oi.Suffix = in["suffix"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsLocalDnsResolution1531(d []interface{}) edpt.SlbTemplateDnsLocalDnsResolution1531 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsLocalDnsResolution1531
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostListCfg = getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg1532(in["host_list_cfg"].([]interface{}))
		ret.LocalResolverCfg = getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533(in["local_resolver_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg1532(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg1532 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg1532, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg1532
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533
		oi.LocalResolver = in["local_resolver"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsNegativeDnsCache1534(d []interface{}) edpt.SlbTemplateDnsNegativeDnsCache1534 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsNegativeDnsCache1534
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableNegativeDnsCache = in["enable_negative_dns_cache"].(int)
		ret.BypassQueryThreshold = in["bypass_query_threshold"].(int)
		ret.MaxNegativeCacheTtl = in["max_negative_cache_ttl"].(int)
		ret.CacheNonValid = in["cache_non_valid"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbTemplateDnsQueryClassFilter1535(d []interface{}) edpt.SlbTemplateDnsQueryClassFilter1535 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsQueryClassFilter1535
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryClassAction = in["query_class_action"].(string)
		ret.QueryClass = getSliceSlbTemplateDnsQueryClassFilterQueryClass1536(in["query_class"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsQueryClassFilterQueryClass1536(d []interface{}) []edpt.SlbTemplateDnsQueryClassFilterQueryClass1536 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryClassFilterQueryClass1536, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryClassFilterQueryClass1536
		oi.StrQueryClass = in["str_query_class"].(string)
		oi.NumQueryClass = in["num_query_class"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsQueryTypeFilter1537(d []interface{}) edpt.SlbTemplateDnsQueryTypeFilter1537 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsQueryTypeFilter1537
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryTypeAction = in["query_type_action"].(string)
		ret.QueryType = getSliceSlbTemplateDnsQueryTypeFilterQueryType1538(in["query_type"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsQueryTypeFilterQueryType1538(d []interface{}) []edpt.SlbTemplateDnsQueryTypeFilterQueryType1538 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryTypeFilterQueryType1538, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryTypeFilterQueryType1538
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolution1539(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolution1539 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolution1539
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostListCfg = getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg1540(in["host_list_cfg"].([]interface{}))
		ret.CsubnetRetry = in["csubnet_retry"].(int)
		ret.NsCacheLookup = in["ns_cache_lookup"].(string)
		ret.NsLongestMatch = in["ns_longest_match"].(string)
		ret.UseServiceGroupResponse = in["use_service_group_response"].(string)
		ret.Ipv4NatPool = in["ipv4_nat_pool"].(string)
		ret.Ipv6NatPool = in["ipv6_nat_pool"].(string)
		ret.RetriesPerLevel = in["retries_per_level"].(int)
		ret.ParallelQueries = in["parallel_queries"].(int)
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
		ret.EdnsUdpSize = in["edns_udp_size"].(int)
		ret.MaxSignatureValidationAttempts = in["max_signature_validation_attempts"].(int)
		ret.MaxSignatureValidationFailures = in["max_signature_validation_failures"].(int)
		ret.MaxKeyDigestValidationFailures = in["max_key_digest_validation_failures"].(int)
		//omit uuid
		ret.LookupOrder = getObjectSlbTemplateDnsRecursiveDnsResolutionLookupOrder1541(in["lookup_order"].([]interface{}))
		ret.GatewayHealthCheck = getObjectSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1543(in["gateway_health_check"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg1540(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg1540 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg1540, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg1540
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionLookupOrder1541(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder1541 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder1541
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryType = getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542(in["query_type"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		oi.Order = in["order"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1543(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1543 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1543
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

func getObjectSlbTemplateDnsResponseRateLimiting1544(d []interface{}) edpt.SlbTemplateDnsResponseRateLimiting1544 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsResponseRateLimiting1544
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ResponseRate = in["response_rate"].(int)
		ret.NxResponseRate = in["nx_response_rate"].(int)
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
		ret.RrlClassListList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList1545(in["rrl_class_list_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList1545(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList1545 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList1545, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList1545
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.LidList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546(in["lid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546
		oi.Lidnum = in["lidnum"].(int)
		oi.LidResponseRate = in["lid_response_rate"].(int)
		oi.LidSlipRate = in["lid_slip_rate"].(int)
		oi.LidNxResponseRate = in["lid_nx_response_rate"].(int)
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

func getObjectSlbTemplateDnsUdpRetransmit1547(d []interface{}) edpt.SlbTemplateDnsUdpRetransmit1547 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsUdpRetransmit1547
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
	ret.Inst.CacheHitcountEnable = d.Get("cache_hitcount_enable").(int)
	ret.Inst.CacheRecordServingPolicy = d.Get("cache_record_serving_policy").(string)
	ret.Inst.CacheTtlAdjustmentEnable = d.Get("cache_ttl_adjustment_enable").(int)
	ret.Inst.CategoryLookupBypass = d.Get("category_lookup_bypass").(string)
	ret.Inst.CategoryLookupList = getSliceSlbTemplateDnsCategoryLookupList(d.Get("category_lookup_list").([]interface{}))
	ret.Inst.CategoryLookupOnlineLookup = d.Get("category_lookup_online_lookup").(int)
	ret.Inst.ClassList = getObjectSlbTemplateDnsClassList1524(d.Get("class_list").([]interface{}))
	ret.Inst.DefaultPolicy = d.Get("default_policy").(string)
	ret.Inst.DisableDnsTemplate = d.Get("disable_dns_template").(int)
	ret.Inst.DisableRaCachedResp = d.Get("disable_ra_cached_resp").(int)
	ret.Inst.DisableRpzAttachSoa = d.Get("disable_rpz_attach_soa").(int)
	ret.Inst.DnsCookieCachePolicy = d.Get("dns_cookie_cache_policy").(string)
	ret.Inst.DnsLogging = d.Get("dns_logging").(string)
	ret.Inst.Dns64 = getObjectSlbTemplateDnsDns641527(d.Get("dns64").([]interface{}))
	ret.Inst.DnssecServiceGroup = d.Get("dnssec_service_group").(string)
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.EnableCacheSharing = d.Get("enable_cache_sharing").(int)
	ret.Inst.Forward = d.Get("forward").(string)
	ret.Inst.InsertIpv4 = d.Get("insert_ipv4").(int)
	ret.Inst.InsertIpv6 = d.Get("insert_ipv6").(int)
	ret.Inst.LabelCountFilter = getObjectSlbTemplateDnsLabelCountFilter1528(d.Get("label_count_filter").([]interface{}))
	ret.Inst.LabelLengthFilter = getObjectSlbTemplateDnsLabelLengthFilter1529(d.Get("label_length_filter").([]interface{}))
	ret.Inst.LocalDnsResolution = getObjectSlbTemplateDnsLocalDnsResolution1531(d.Get("local_dns_resolution").([]interface{}))
	ret.Inst.MaxCacheEntrySize = d.Get("max_cache_entry_size").(int)
	ret.Inst.MaxCacheSize = d.Get("max_cache_size").(int)
	ret.Inst.MaxQueryLength = d.Get("max_query_length").(int)
	ret.Inst.MaxUdpSize = d.Get("max_udp_size").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NegativeDnsCache = getObjectSlbTemplateDnsNegativeDnsCache1534(d.Get("negative_dns_cache").([]interface{}))
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.QpsLogHigh = d.Get("qps_log_high").(int)
	ret.Inst.QpsLogLow = d.Get("qps_log_low").(int)
	ret.Inst.QpsThresholdLog = d.Get("qps_threshold_log").(int)
	ret.Inst.QueryClassFilter = getObjectSlbTemplateDnsQueryClassFilter1535(d.Get("query_class_filter").([]interface{}))
	ret.Inst.QueryIdSwitch = d.Get("query_id_switch").(int)
	ret.Inst.QueryTypeFilter = getObjectSlbTemplateDnsQueryTypeFilter1537(d.Get("query_type_filter").([]interface{}))
	ret.Inst.RecursiveDnsResolution = getObjectSlbTemplateDnsRecursiveDnsResolution1539(d.Get("recursive_dns_resolution").([]interface{}))
	ret.Inst.RedirectToTcpPort = d.Get("redirect_to_tcp_port").(int)
	ret.Inst.RemoveAaFlag = d.Get("remove_aa_flag").(int)
	ret.Inst.RemoveCsubnet = d.Get("remove_csubnet").(int)
	ret.Inst.RemovePaddingToServer = d.Get("remove_padding_to_server").(int)
	ret.Inst.ResponseRateLimiting = getObjectSlbTemplateDnsResponseRateLimiting1544(d.Get("response_rate_limiting").([]interface{}))
	ret.Inst.RpzList = getSliceSlbTemplateDnsRpzList(d.Get("rpz_list").([]interface{}))
	ret.Inst.TldFilterLogEnable = d.Get("tld_filter_log_enable").(int)
	ret.Inst.TldFilterWhiteList = d.Get("tld_filter_white_list").(string)
	ret.Inst.UdpRetransmit = getObjectSlbTemplateDnsUdpRetransmit1547(d.Get("udp_retransmit").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
