package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSlbTemplateDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns`: DNS template\n\n",
		CreateContext: resourceSlbTemplateDnsCreate,
		UpdateContext: resourceSlbTemplateDnsUpdate,
		ReadContext:   resourceSlbTemplateDnsRead,
		DeleteContext: resourceSlbTemplateDnsDelete,
		Schema: map[string]*schema.Schema{
			"add_padding_to_client": {
				Type: schema.TypeString, Optional: true, Description: "'block-length': Block-Length Padding; 'random-block-length': Random-Block-Length Padding;",
				ValidateFunc: validation.StringInSlice([]string{"block-length", "random-block-length"}, false),
			},
			"cache_record_serving_policy": {
				Type: schema.TypeString, Optional: true, Description: "'global': Follow global cofiguration (Default); 'no-change': No change in record order; 'round-robin': Round-robin;",
				ValidateFunc: validation.StringInSlice([]string{"global", "no-change", "round-robin"}, false),
			},
			"class_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify a class list name",
							ValidateFunc: validation.StringLenBetween(1, 63),
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
										ValidateFunc: validation.IntBetween(1, 1023),
									},
									"conn_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
										ValidateFunc: validation.IntBetween(1, 2147483647),
									},
									"per": {
										Type: schema.TypeInt, Optional: true, Description: "Per (Number of 100ms)",
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"over_limit_action": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action when exceeds limit",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"action_value": {
										Type: schema.TypeString, Optional: true, Description: "'dns-cache-disable': Disable DNS cache when it exceeds limit; 'dns-cache-enable': Enable DNS cache when it exceeds limit; 'forward': Forward the traffic even it exceeds limit;",
										ValidateFunc: validation.StringInSlice([]string{"dns-cache-disable", "dns-cache-enable", "forward"}, false),
									},
									"lockout": {
										Type: schema.TypeInt, Optional: true, Description: "Don't accept any new connection for certain time (Lockout duration in minutes)",
										ValidateFunc: validation.IntBetween(1, 1023),
									},
									"log": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log a message",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"log_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Log interval (minute, by default system will log every over limit instance)",
										ValidateFunc: validation.IntBetween(1, 255),
									},
									"dns": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cache_action": {
													Type: schema.TypeString, Optional: true, Default: "cache-disable", Description: "'cache-disable': Disable dns cache; 'cache-enable': Enable dns cache;",
													ValidateFunc: validation.StringInSlice([]string{"cache-disable", "cache-enable"}, false),
												},
												"ttl": {
													Type: schema.TypeInt, Optional: true, Description: "TTL for cache entry (TTL in seconds)",
													ValidateFunc: validation.IntBetween(1, 65535),
												},
												"weight": {
													Type: schema.TypeInt, Optional: true, Description: "Weight for cache entry",
													ValidateFunc: validation.IntBetween(1, 7),
												},
												"honor_server_response_ttl": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Honor the server reponse TTL",
													ValidateFunc: validation.IntBetween(0, 1),
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
								},
							},
						},
					},
				},
			},
			"default_policy": {
				Type: schema.TypeString, Optional: true, Default: "nocache", Description: "'nocache': Cache disable; 'cache': Cache enable;",
				ValidateFunc: validation.StringInSlice([]string{"nocache", "cache"}, false),
			},
			"disable_dns_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS template",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"disable_ra_cached_resp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS recursive available flag in cached response",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"disable_rpz_attach_soa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable attaching SOA due to RPZ",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"dns_logging": {
				Type: schema.TypeString, Optional: true, Description: "dns logging template (DNS Logging template name)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"dnssec_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Use different service group if DNSSEC DO bit set (Service Group Name)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop the malformed query",
				ValidateFunc:  validation.IntBetween(0, 1),
				ConflictsWith: []string{"forward"},
			},
			"enable_cache_sharing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DNS cache sharing",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"forward": {
				Type: schema.TypeString, Optional: true, Description: "Forward to service group (Service group name)",
				ValidateFunc:  validation.StringLenBetween(1, 127),
				ConflictsWith: []string{"drop"},
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
										Type: schema.TypeString, Optional: true, Description: "Hostnames class-list name (ac type)",
										ValidateFunc: validation.StringLenBetween(1, 63),
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
										ValidateFunc: validation.IsIPv4Address,
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
				ValidateFunc: validation.IntBetween(1, 4096),
			},
			"max_cache_size": {
				Type: schema.TypeInt, Optional: true, Description: "Define maximum cache size (Maximum cache entry per VIP)",
			},
			"max_query_length": {
				Type: schema.TypeInt, Optional: true, Description: "Define Maximum DNS Query Length, default is unlimited (Specify Maximum Length)",
				ValidateFunc: validation.IntBetween(1, 4095),
			},
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "DNS Template Name",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Period in minutes",
				ValidateFunc: validation.IntBetween(1, 10000),
			},
			"query_class_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_class_action": {
							Type: schema.TypeString, Optional: true, Description: "'allow': Allow only certain DNS query classes; 'deny': Deny only certain DNS query classes;",
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),
						},
						"query_class": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_class": {
										Type: schema.TypeString, Optional: true, Description: "'INTERNET': INTERNET query class; 'CHAOS': CHAOS query class; 'HESIOD': HESIOD query class; 'NONE': NONE query class; 'ANY': ANY query class;",
										ValidateFunc: validation.StringInSlice([]string{"INTERNET", "CHAOS", "HESIOD", "NONE", "ANY"}, false),
									},
									"num_query_class": {
										Type: schema.TypeInt, Optional: true, Description: "Other query class value",
										ValidateFunc: validation.IntBetween(1, 65535),
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
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"query_type_filter": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_type_action": {
							Type: schema.TypeString, Optional: true, Description: "'allow': Allow only certain DNS query types; 'deny': Deny only certain DNS query types;",
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),
						},
						"query_type": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_type": {
										Type: schema.TypeString, Optional: true, Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;",
										ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CNAME", "MX", "NS", "SRV", "PTR", "SOA", "TXT", "ANY"}, false),
									},
									"num_query_type": {
										Type: schema.TypeInt, Optional: true, Description: "Other record type value",
										ValidateFunc: validation.IntBetween(1, 65535),
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
										Type: schema.TypeString, Optional: true, Description: "Hostnames class-list name (ac type)",
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
								},
							},
						},
						"ns_cache_lookup": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Disable NS Cache Lookup; 'enabled': Enable NS Cache Lookup;",
							ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
						},
						"use_service_group_response": {
							Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Start Recursive Resolver if Server response doesnt have final answer; 'enabled': Forward Backend Server response to client and dont start recursive resolver;",
							ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
						},
						"ipv4_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Source NAT pool or pool group",
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"ipv6_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Source NAT pool or pool group",
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"retries_per_level": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Number of DNS query retries at each server level before closing client connection, default 6",
							ValidateFunc: validation.IntBetween(1, 6),
						},
						"full_response": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Serve all records (authority and additional) when applicable",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"max_trials": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Total number of times to try DNS query to server before closing client connection, default 0",
							ValidateFunc: validation.IntBetween(0, 255),
						},
						"use_client_qid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use client side query id for recursive query",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"redirect_to_tcp_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Direct the client to retry with TCP for DNS UDP request",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"remove_aa_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Make answers created from cache non-authoritative",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"remove_edns_csubnet_to_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove EDNS(0) client subnet from client queries",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"remove_padding_to_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove EDNS(0) padding to server",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"response_rate_limiting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_rate": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
							ValidateFunc: validation.IntBetween(1, 1000),
						},
						"filter_response_rate": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Maximum allowed request rate for the filter. This should match average traffic. (default 10 per seconds)",
							ValidateFunc: validation.IntBetween(1, 1000),
						},
						"slip_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
							ValidateFunc: validation.IntBetween(2, 10),
						},
						"window": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
							ValidateFunc: validation.IntBetween(1, 60),
						},
						"enable_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
							ValidateFunc: validation.StringInSlice([]string{"log-only", "rate-limit", "whitelist"}, false),
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
										ValidateFunc: validation.StringLenBetween(1, 63),
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
										ValidateFunc: validation.StringLenBetween(1, 127),
									},
									"lid_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lidnum": {
													Type: schema.TypeInt, Required: true, Description: "Specify a limit ID",
													ValidateFunc: validation.IntBetween(1, 1023),
												},
												"lid_response_rate": {
													Type: schema.TypeInt, Optional: true, Default: 5, Description: "Responses exceeding this rate within the window will be dropped (default 5 per second)",
													ValidateFunc: validation.IntBetween(1, 1000),
												},
												"lid_slip_rate": {
													Type: schema.TypeInt, Optional: true, Description: "Every n'th response that would be rate-limited will be let through instead",
													ValidateFunc: validation.IntBetween(2, 10),
												},
												"lid_window": {
													Type: schema.TypeInt, Optional: true, Default: 1, Description: "Rate-Limiting Interval in Seconds (default is one)",
													ValidateFunc: validation.IntBetween(1, 60),
												},
												"lid_enable_log": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
													ValidateFunc: validation.IntBetween(0, 1),
												},
												"lid_action": {
													Type: schema.TypeString, Optional: true, Default: "rate-limit", Description: "'log-only': Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit': Rate-Limit based on configuration (Default); 'whitelist': Whitelist, disable rate-limiting;",
													ValidateFunc: validation.StringInSlice([]string{"log-only", "rate-limit", "whitelist"}, false),
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
												"user_tag": {
													Type: schema.TypeString, Optional: true, Description: "Customized tag",
													ValidateFunc: validation.StringLenBetween(1, 127),
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
							ValidateFunc: validation.IntBetween(1, 8),
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify a Response Policy Zone name",
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
							ValidateFunc: validation.StringLenBetween(1, 127),
						},
						"logging": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log RPZ triggered action",
										ValidateFunc: validation.IntBetween(0, 1),
									},
									"rpz_action": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str_rpz_action": {
													Type: schema.TypeString, Optional: true, Description: "'drop': Log RPZ due to drop action; 'pass-thru': Log RPZ due to pass-thru action; 'nxdomain': Log RPZ due to nxdomain action; 'nodata': Log RPZ due to nodata action; 'tcp-only': Log RPZ due to tcp-only action; 'local-data': Log RPZ due to local-data action;",
													ValidateFunc: validation.StringInSlice([]string{"drop", "pass-thru", "nxdomain", "nodata", "tcp-only", "local-data"}, false),
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
							ValidateFunc: validation.IntBetween(1, 400),
						},
						"max_trials": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Total number of times to try DNS query to server before closing client connection, default 3",
							ValidateFunc: validation.IntBetween(1, 5),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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

func getObjectSlbTemplateDnsClassList(d []interface{}) edpt.SlbTemplateDnsClassList {
	count := len(d)
	var ret edpt.SlbTemplateDnsClassList
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		//omit uuid
		ret.LidList = getSliceSlbTemplateDnsClassListLidList(in["lid_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsClassListLidList(d []interface{}) []edpt.SlbTemplateDnsClassListLidList {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsClassListLidList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsClassListLidList
		oi.Lidnum = in["lidnum"].(int)
		oi.ConnRateLimit = in["conn_rate_limit"].(int)
		oi.Per = in["per"].(int)
		oi.OverLimitAction = in["over_limit_action"].(int)
		oi.ActionValue = in["action_value"].(string)
		oi.Lockout = in["lockout"].(int)
		oi.Log = in["log"].(int)
		oi.LogInterval = in["log_interval"].(int)
		oi.Dns = getObjectSlbTemplateDnsClassListLidListDns(in["dns"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsClassListLidListDns(d []interface{}) edpt.SlbTemplateDnsClassListLidListDns {
	count := len(d)
	var ret edpt.SlbTemplateDnsClassListLidListDns
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheAction = in["cache_action"].(string)
		ret.Ttl = in["ttl"].(int)
		ret.Weight = in["weight"].(int)
		ret.HonorServerResponseTtl = in["honor_server_response_ttl"].(int)
	}
	return ret
}

func getObjectSlbTemplateDnsLocalDnsResolution(d []interface{}) edpt.SlbTemplateDnsLocalDnsResolution {
	count := len(d)
	var ret edpt.SlbTemplateDnsLocalDnsResolution
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.HostListCfg = getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg(in["host_list_cfg"].([]interface{}))
		ret.LocalResolverCfg = getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg(in["local_resolver_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionHostListCfg(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionHostListCfg
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsLocalDnsResolutionLocalResolverCfg(d []interface{}) []edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLocalDnsResolutionLocalResolverCfg
		oi.LocalResolver = in["local_resolver"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsQueryClassFilter(d []interface{}) edpt.SlbTemplateDnsQueryClassFilter {
	count := len(d)
	var ret edpt.SlbTemplateDnsQueryClassFilter
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryClassAction = in["query_class_action"].(string)
		ret.QueryClass = getSliceSlbTemplateDnsQueryClassFilterQueryClass(in["query_class"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsQueryClassFilterQueryClass(d []interface{}) []edpt.SlbTemplateDnsQueryClassFilterQueryClass {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryClassFilterQueryClass, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryClassFilterQueryClass
		oi.StrQueryClass = in["str_query_class"].(string)
		oi.NumQueryClass = in["num_query_class"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsQueryTypeFilter(d []interface{}) edpt.SlbTemplateDnsQueryTypeFilter {
	count := len(d)
	var ret edpt.SlbTemplateDnsQueryTypeFilter
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryTypeAction = in["query_type_action"].(string)
		ret.QueryType = getSliceSlbTemplateDnsQueryTypeFilterQueryType(in["query_type"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsQueryTypeFilterQueryType(d []interface{}) []edpt.SlbTemplateDnsQueryTypeFilterQueryType {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryTypeFilterQueryType, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryTypeFilterQueryType
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolution(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolution {
	count := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolution
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.HostListCfg = getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg(in["host_list_cfg"].([]interface{}))
		ret.NsCacheLookup = in["ns_cache_lookup"].(string)
		ret.UseServiceGroupResponse = in["use_service_group_response"].(string)
		ret.Ipv4NatPool = in["ipv4_nat_pool"].(string)
		ret.Ipv6NatPool = in["ipv6_nat_pool"].(string)
		ret.RetriesPerLevel = in["retries_per_level"].(int)
		ret.FullResponse = in["full_response"].(int)
		ret.MaxTrials = in["max_trials"].(int)
		ret.UseClientQid = in["use_client_qid"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsResponseRateLimiting(d []interface{}) edpt.SlbTemplateDnsResponseRateLimiting {
	count := len(d)
	var ret edpt.SlbTemplateDnsResponseRateLimiting
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.ResponseRate = in["response_rate"].(int)
		ret.FilterResponseRate = in["filter_response_rate"].(int)
		ret.SlipRate = in["slip_rate"].(int)
		ret.Window = in["window"].(int)
		ret.EnableLog = in["enable_log"].(int)
		ret.Action = in["action"].(string)
		//omit uuid
		ret.RrlClassListList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList(in["rrl_class_list_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListList(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.LidList = getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList(in["lid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsResponseRateLimitingRrlClassListListLidList(d []interface{}) []edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsResponseRateLimitingRrlClassListListLidList
		oi.Lidnum = in["lidnum"].(int)
		oi.LidResponseRate = in["lid_response_rate"].(int)
		oi.LidSlipRate = in["lid_slip_rate"].(int)
		oi.LidWindow = in["lid_window"].(int)
		oi.LidEnableLog = in["lid_enable_log"].(int)
		oi.LidAction = in["lid_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDnsRpzList(d []interface{}) []edpt.SlbTemplateDnsRpzList {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsRpzList, 0, count)
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
	count := len(d)
	var ret edpt.SlbTemplateDnsRpzListLogging
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.RpzAction = getSliceSlbTemplateDnsRpzListLoggingRpzAction(in["rpz_action"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRpzListLoggingRpzAction(d []interface{}) []edpt.SlbTemplateDnsRpzListLoggingRpzAction {
	count := len(d)
	ret := make([]edpt.SlbTemplateDnsRpzListLoggingRpzAction, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRpzListLoggingRpzAction
		oi.StrRpzAction = in["str_rpz_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsUdpRetransmit(d []interface{}) edpt.SlbTemplateDnsUdpRetransmit {
	count := len(d)
	var ret edpt.SlbTemplateDnsUdpRetransmit
	if count > 0 {
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
	ret.Inst.ClassList = getObjectSlbTemplateDnsClassList(d.Get("class_list").([]interface{}))
	ret.Inst.DefaultPolicy = d.Get("default_policy").(string)
	ret.Inst.DisableDnsTemplate = d.Get("disable_dns_template").(int)
	ret.Inst.DisableRaCachedResp = d.Get("disable_ra_cached_resp").(int)
	ret.Inst.DisableRpzAttachSoa = d.Get("disable_rpz_attach_soa").(int)
	ret.Inst.DnsLogging = d.Get("dns_logging").(string)
	ret.Inst.DnssecServiceGroup = d.Get("dnssec_service_group").(string)
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.EnableCacheSharing = d.Get("enable_cache_sharing").(int)
	ret.Inst.Forward = d.Get("forward").(string)
	ret.Inst.LocalDnsResolution = getObjectSlbTemplateDnsLocalDnsResolution(d.Get("local_dns_resolution").([]interface{}))
	ret.Inst.MaxCacheEntrySize = d.Get("max_cache_entry_size").(int)
	ret.Inst.MaxCacheSize = d.Get("max_cache_size").(int)
	ret.Inst.MaxQueryLength = d.Get("max_query_length").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.QueryClassFilter = getObjectSlbTemplateDnsQueryClassFilter(d.Get("query_class_filter").([]interface{}))
	ret.Inst.QueryIdSwitch = d.Get("query_id_switch").(int)
	ret.Inst.QueryTypeFilter = getObjectSlbTemplateDnsQueryTypeFilter(d.Get("query_type_filter").([]interface{}))
	ret.Inst.RecursiveDnsResolution = getObjectSlbTemplateDnsRecursiveDnsResolution(d.Get("recursive_dns_resolution").([]interface{}))
	ret.Inst.RedirectToTcpPort = d.Get("redirect_to_tcp_port").(int)
	ret.Inst.RemoveAaFlag = d.Get("remove_aa_flag").(int)
	ret.Inst.RemoveEdnsCsubnetToServer = d.Get("remove_edns_csubnet_to_server").(int)
	ret.Inst.RemovePaddingToServer = d.Get("remove_padding_to_server").(int)
	ret.Inst.ResponseRateLimiting = getObjectSlbTemplateDnsResponseRateLimiting(d.Get("response_rate_limiting").([]interface{}))
	ret.Inst.RpzList = getSliceSlbTemplateDnsRpzList(d.Get("rpz_list").([]interface{}))
	ret.Inst.UdpRetransmit = getObjectSlbTemplateDnsUdpRetransmit(d.Get("udp_retransmit").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
