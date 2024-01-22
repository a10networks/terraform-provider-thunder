package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_dns`: DNS template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateDnsCreate,
		UpdateContext: resourceDdosTemplateDnsUpdate,
		ReadContext:   resourceDdosTemplateDnsRead,
		DeleteContext: resourceDdosTemplateDnsDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default action); 'reset': Send Client RST for TCP connections;",
			},
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
					},
				},
			},
			"dns_any_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop DNS queries of Type ANY",
			},
			"dns_auth_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_auth": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DNS authentication",
						},
						"dns_auth_type": {
							Type: schema.TypeString, Optional: true, Description: "'udp': Drop DNS request and monitor client retry; 'force-tcp': Force DNS request over TCP;",
						},
						"udp_timeout_val_only": {
							Type: schema.TypeInt, Optional: true, Description: "UDP authentication timeout in seconds",
						},
						"udp_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "UDP authentication timeout in seconds",
						},
						"min_retry_gap": {
							Type: schema.TypeInt, Optional: true, Description: "Optional minimum sec gap in between 2 dns-udp packets for auth to pass, unit is specified by min-retry-gap-interval",
						},
						"min_retry_gap_interval": {
							Type: schema.TypeString, Optional: true, Default: "1sec", Description: "'100ms': 100ms; '1sec': 1sec;",
						},
						"with_udp_auth": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Monitor client retry",
						},
						"force_tcp_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP authentication timeout in seconds",
						},
						"force_tcp_min_retry_gap": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum sec gap in between 2 dns-udp packets for auth to pass",
						},
						"force_tcp_ignore_client_source_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow client to retransmit DNS request using different source port during udp-auth (supported in asymmetric mode only)",
						},
					},
				},
			},
			"dns_request_rate_limit": {
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
					},
				},
			},
			"domain_group_name": {
				Type: schema.TypeString, Optional: true, Description: "Apply a domain-group to the DNS template",
			},
			"domain_group_rate_exceed_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop the query (default); 'tunnel-encap-packet': Encapsulate the query and send on a tunnel;",
			},
			"domain_group_rate_per_service": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable per service domain rate checking",
			},
			"encap_template": {
				Type: schema.TypeString, Optional: true, Description: "DDOS encap template to sepcify the tunnel endpoint",
			},
			"fqdn_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_fqdn_rate_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DNS Rate limiting on the basis of FQDN",
						},
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
						"by": {
							Type: schema.TypeString, Optional: true, Description: "'domain-name': Domain Name; 'src-ip': Source IP address; 'both': Use both Domain Name and Source IP address for rate-limiting;",
						},
						"fqdn_rate_suffix_by": {
							Type: schema.TypeInt, Optional: true, Description: "Number of suffixes",
						},
					},
				},
			},
			"fqdn_label_count": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of length of FQDN labels",
			},
			"fqdn_label_len_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn_label_length": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Maximum FQDN label length",
						},
						"label_length": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum length of FQDN label",
						},
						"fqdn_label_suffix": {
							Type: schema.TypeInt, Optional: true, Description: "Number of suffixes",
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
			"nxdomain_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_nxdomain_rate_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DNS NXDOMAIN Rate Limiting (SRC support only)",
						},
						"dns_nxdomain_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Limiting rate",
						},
						"dns_nxdomain_rate_limit_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop queries if rate is exceeded; 'black-list': Black-List source if rate is exceeded;",
						},
					},
				},
			},
			"on_no_match": {
				Type: schema.TypeString, Optional: true, Default: "deny", Description: "'permit': permit; 'deny': deny (default);",
			},
			"query_rate_threshold_for_cache_serving": {
				Type: schema.TypeInt, Optional: true, Description: "This is for DNS cache mode only, it sets a DNS query rate threshold such that queries under the rate threshold would be forward",
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
func resourceDdosTemplateDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateDnsAllowQueryClass(d []interface{}) edpt.DdosTemplateDnsAllowQueryClass {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsAllowQueryClass
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllowInternetQueryClass = in["allow_internet_query_class"].(int)
		ret.AllowCsnetQueryClass = in["allow_csnet_query_class"].(int)
		ret.AllowChaosQueryClass = in["allow_chaos_query_class"].(int)
		ret.AllowHesiodQueryClass = in["allow_hesiod_query_class"].(int)
		ret.AllowNoneQueryClass = in["allow_none_query_class"].(int)
		ret.AllowAnyQueryClass = in["allow_any_query_class"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsAllowRecordType(d []interface{}) edpt.DdosTemplateDnsAllowRecordType {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsAllowRecordType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllowAType = in["allow_a_type"].(int)
		ret.AllowAaaaType = in["allow_aaaa_type"].(int)
		ret.AllowCnameType = in["allow_cname_type"].(int)
		ret.AllowMxType = in["allow_mx_type"].(int)
		ret.AllowNsType = in["allow_ns_type"].(int)
		ret.AllowSrvType = in["allow_srv_type"].(int)
		ret.RecordNumCfg = getSliceDdosTemplateDnsAllowRecordTypeRecordNumCfg(in["record_num_cfg"].([]interface{}))
	}
	return ret
}

func getSliceDdosTemplateDnsAllowRecordTypeRecordNumCfg(d []interface{}) []edpt.DdosTemplateDnsAllowRecordTypeRecordNumCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateDnsAllowRecordTypeRecordNumCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateDnsAllowRecordTypeRecordNumCfg
		oi.AllowNumType = in["allow_num_type"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsAuthCfg(d []interface{}) edpt.DdosTemplateDnsDnsAuthCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsAuthCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsAuth = in["dns_auth"].(int)
		ret.DnsAuthType = in["dns_auth_type"].(string)
		ret.UdpTimeoutValOnly = in["udp_timeout_val_only"].(int)
		ret.UdpTimeout = in["udp_timeout"].(int)
		ret.MinRetryGap = in["min_retry_gap"].(int)
		ret.MinRetryGapInterval = in["min_retry_gap_interval"].(string)
		ret.WithUdpAuth = in["with_udp_auth"].(int)
		ret.ForceTcpTimeout = in["force_tcp_timeout"].(int)
		ret.ForceTcpMinRetryGap = in["force_tcp_min_retry_gap"].(int)
		ret.ForceTcpIgnoreClientSourcePort = in["force_tcp_ignore_client_source_port"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimit(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimit {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = getObjectDdosTemplateDnsDnsRequestRateLimitType(in["type"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitType(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitType {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ACfg = getObjectDdosTemplateDnsDnsRequestRateLimitTypeACfg(in["a_cfg"].([]interface{}))
		ret.AaaaCfg = getObjectDdosTemplateDnsDnsRequestRateLimitTypeAaaaCfg(in["aaaa_cfg"].([]interface{}))
		ret.CnameCfg = getObjectDdosTemplateDnsDnsRequestRateLimitTypeCnameCfg(in["cname_cfg"].([]interface{}))
		ret.MxCfg = getObjectDdosTemplateDnsDnsRequestRateLimitTypeMxCfg(in["mx_cfg"].([]interface{}))
		ret.NsCfg = getObjectDdosTemplateDnsDnsRequestRateLimitTypeNsCfg(in["ns_cfg"].([]interface{}))
		ret.SrvCfg = getObjectDdosTemplateDnsDnsRequestRateLimitTypeSrvCfg(in["srv_cfg"].([]interface{}))
		ret.DnsTypeCfg = getSliceDdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg(in["dns_type_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitTypeACfg(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitTypeACfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitTypeACfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.A = in["a"].(int)
		ret.DnsARate = in["dns_a_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitTypeAaaaCfg(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitTypeAaaaCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitTypeAaaaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Aaaa = in["aaaa"].(int)
		ret.DnsAaaaRate = in["dns_aaaa_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitTypeCnameCfg(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitTypeCnameCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitTypeCnameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cname = in["cname"].(int)
		ret.DnsCnameRate = in["dns_cname_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitTypeMxCfg(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitTypeMxCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitTypeMxCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mx = in["mx"].(int)
		ret.DnsMxRate = in["dns_mx_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitTypeNsCfg(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitTypeNsCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitTypeNsCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ns = in["ns"].(int)
		ret.DnsNsRate = in["dns_ns_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateDnsDnsRequestRateLimitTypeSrvCfg(d []interface{}) edpt.DdosTemplateDnsDnsRequestRateLimitTypeSrvCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsDnsRequestRateLimitTypeSrvCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Srv = in["srv"].(int)
		ret.DnsSrvRate = in["dns_srv_rate"].(int)
	}
	return ret
}

func getSliceDdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg(d []interface{}) []edpt.DdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg
		oi.DnsRequestType = in["dns_request_type"].(int)
		oi.DnsRequestTypeRate = in["dns_request_type_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateDnsFqdnCfg(d []interface{}) []edpt.DdosTemplateDnsFqdnCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateDnsFqdnCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateDnsFqdnCfg
		oi.DnsFqdnRateLimit = in["dns_fqdn_rate_limit"].(int)
		oi.DnsFqdnRate = in["dns_fqdn_rate"].(int)
		oi.Per = in["per"].(string)
		oi.PerDomainPerSrcIp = in["per_domain_per_src_ip"].(int)
		oi.FqdnRateSuffix = in["fqdn_rate_suffix"].(int)
		oi.FqdnRateLabelCount = in["fqdn_rate_label_count"].(int)
		oi.By = in["by"].(string)
		oi.FqdnRateSuffixBy = in["fqdn_rate_suffix_by"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateDnsFqdnLabelLenCfg(d []interface{}) []edpt.DdosTemplateDnsFqdnLabelLenCfg {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateDnsFqdnLabelLenCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateDnsFqdnLabelLenCfg
		oi.FqdnLabelLength = in["fqdn_label_length"].(int)
		oi.LabelLength = in["label_length"].(int)
		oi.FqdnLabelSuffix = in["fqdn_label_suffix"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateDnsMalformedQueryCheck295(d []interface{}) edpt.DdosTemplateDnsMalformedQueryCheck295 {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsMalformedQueryCheck295
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ValidationType = in["validation_type"].(string)
		ret.NonQueryOpcodeCheck = in["non_query_opcode_check"].(string)
		ret.SkipMultiPacketCheck = in["skip_multi_packet_check"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosTemplateDnsMultiPuThresholdDistribution(d []interface{}) edpt.DdosTemplateDnsMultiPuThresholdDistribution {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsMultiPuThresholdDistribution
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiPuThresholdDistributionValue = in["multi_pu_threshold_distribution_value"].(int)
		ret.MultiPuThresholdDistributionDisable = in["multi_pu_threshold_distribution_disable"].(string)
	}
	return ret
}

func getObjectDdosTemplateDnsNxdomainCfg(d []interface{}) edpt.DdosTemplateDnsNxdomainCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsNxdomainCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsNxdomainRateLimit = in["dns_nxdomain_rate_limit"].(int)
		ret.DnsNxdomainRate = in["dns_nxdomain_rate"].(int)
		ret.DnsNxdomainRateLimitAction = in["dns_nxdomain_rate_limit_action"].(string)
	}
	return ret
}

func getObjectDdosTemplateDnsSymtimeoutCfg(d []interface{}) edpt.DdosTemplateDnsSymtimeoutCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateDnsSymtimeoutCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SymTimeout = in["sym_timeout"].(int)
		ret.SymTimeoutValue = in["sym_timeout_value"].(int)
	}
	return ret
}

func dataToEndpointDdosTemplateDns(d *schema.ResourceData) edpt.DdosTemplateDns {
	var ret edpt.DdosTemplateDns
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AllowQueryClass = getObjectDdosTemplateDnsAllowQueryClass(d.Get("allow_query_class").([]interface{}))
	ret.Inst.AllowRecordType = getObjectDdosTemplateDnsAllowRecordType(d.Get("allow_record_type").([]interface{}))
	ret.Inst.DnsAnyCheck = d.Get("dns_any_check").(int)
	ret.Inst.DnsAuthCfg = getObjectDdosTemplateDnsDnsAuthCfg(d.Get("dns_auth_cfg").([]interface{}))
	ret.Inst.DnsRequestRateLimit = getObjectDdosTemplateDnsDnsRequestRateLimit(d.Get("dns_request_rate_limit").([]interface{}))
	ret.Inst.DomainGroupName = d.Get("domain_group_name").(string)
	ret.Inst.DomainGroupRateExceedAction = d.Get("domain_group_rate_exceed_action").(string)
	ret.Inst.DomainGroupRatePerService = d.Get("domain_group_rate_per_service").(int)
	ret.Inst.EncapTemplate = d.Get("encap_template").(string)
	ret.Inst.FqdnCfg = getSliceDdosTemplateDnsFqdnCfg(d.Get("fqdn_cfg").([]interface{}))
	ret.Inst.FqdnLabelCount = d.Get("fqdn_label_count").(int)
	ret.Inst.FqdnLabelLenCfg = getSliceDdosTemplateDnsFqdnLabelLenCfg(d.Get("fqdn_label_len_cfg").([]interface{}))
	ret.Inst.MalformedQueryCheck = getObjectDdosTemplateDnsMalformedQueryCheck295(d.Get("malformed_query_check").([]interface{}))
	ret.Inst.MultiPuThresholdDistribution = getObjectDdosTemplateDnsMultiPuThresholdDistribution(d.Get("multi_pu_threshold_distribution").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NxdomainCfg = getObjectDdosTemplateDnsNxdomainCfg(d.Get("nxdomain_cfg").([]interface{}))
	ret.Inst.OnNoMatch = d.Get("on_no_match").(string)
	ret.Inst.QueryRateThresholdForCacheServing = d.Get("query_rate_threshold_for_cache_serving").(int)
	ret.Inst.SymtimeoutCfg = getObjectDdosTemplateDnsSymtimeoutCfg(d.Get("symtimeout_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
