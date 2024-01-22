package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging`: Logging Template\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingCreate,
		UpdateContext: resourceFwTemplateLoggingUpdate,
		ReadContext:   resourceFwTemplateLoggingRead,
		DeleteContext: resourceFwTemplateLoggingDelete,

		Schema: map[string]*schema.Schema{
			"facility": {
				Type: schema.TypeString, Optional: true, Default: "local0", Description: "'kernel': 0: Kernel; 'user': 1: User-level; 'mail': 2: Mail; 'daemon': 3: System daemons; 'security-authorization': 4: Security/authorization; 'syslog': 5: Syslog internal; 'line-printer': 6: Line printer; 'news': 7: Network news; 'uucp': 8: UUCP subsystem; 'cron': 9: Time-related; 'security-authorization-private': 10: Private security/authorization; 'ftp': 11: FTP; 'ntp': 12: NTP; 'audit': 13: Audit; 'alert': 14: Alert; 'clock': 15: Clock-related; 'local0': 16: Local use 0; 'local1': 17: Local use 1; 'local2': 18: Local use 2; 'local3': 19: Local use 3; 'local4': 20: Local use 4; 'local5': 21: Local use 5; 'local6': 22: Local use 6; 'local7': 23: Local use 7;",
			},
			"format": {
				Type: schema.TypeString, Optional: true, Default: "cef", Description: "'ascii': A10 Text logging format (ASCII); 'cef': Common Event Format for logging (default);",
			},
			"include_dest_fqdn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include destination FQDN string",
			},
			"include_http": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_header": {
										Type: schema.TypeString, Optional: true, Description: "'cookie': Log HTTP Cookie Header; 'referer': Log HTTP Referer Header; 'user-agent': Log HTTP User-Agent Header; 'header1': Log HTTP Header 1; 'header2': Log HTTP Header 2; 'header3': Log HTTP Header 3;",
									},
									"max_length": {
										Type: schema.TypeInt, Optional: true, Default: 100, Description: "Max length for a HTTP request log (Max header length (Default: 100 char))",
									},
									"custom_header_name": {
										Type: schema.TypeString, Optional: true, Description: "Header name",
									},
									"custom_max_length": {
										Type: schema.TypeInt, Optional: true, Default: 100, Description: "Max length for a HTTP request log (Max header length (Default: 100 char))",
									},
								},
							},
						},
						"l4_session_info": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the L4 session information of the HTTP request",
						},
						"method": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the HTTP Request Method",
						},
						"request_number": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP Request Number",
						},
						"file_extension": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP file extension",
						},
					},
				},
			},
			"include_radius_attribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attr": {
										Type: schema.TypeString, Optional: true, Description: "'imei': Include IMEI; 'imsi': Include IMSI; 'msisdn': Include MSISDN; 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;",
									},
									"attr_event": {
										Type: schema.TypeString, Optional: true, Description: "'http-requests': Include in HTTP request logs; 'sessions': Include in session logs; 'limit-policy': Include in limit policy logs;",
									},
								},
							},
						},
						"no_quote": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No quotation marks for RADIUS attributes in logs",
						},
						"framed_ipv6_prefix": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include radius attributes for the prefix",
						},
						"prefix_length": {
							Type: schema.TypeString, Optional: true, Description: "'32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;",
						},
						"insert_if_not_existing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure what string is to be inserted for custom RADIUS attributes",
						},
						"zero_in_custom_attr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert 0000 for standard and custom attributes in log string",
						},
					},
				},
			},
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_requests": {
							Type: schema.TypeString, Optional: true, Description: "'host': Log the HTTP Host Header; 'url': Log the HTTP Request URL;",
						},
					},
				},
			},
			"merged_style": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Merge creation and deletion of session logs to one",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Logging Template Name",
			},
			"resolution": {
				Type: schema.TypeString, Optional: true, Default: "seconds", Description: "'seconds': Logging timestamp resolution in seconds (default); '10-milliseconds': Logging timestamp resolution in 10s of milli-seconds;",
			},
			"rule": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_http_requests": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dest_port": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dest_port_number": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"include_byte_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the byte count of HTTP Request/Response in FW session deletion logs",
												},
											},
										},
									},
									"log_every_http_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log every HTTP request in an HTTP 1.1 session (Default: Log the first HTTP request in a session)",
									},
									"max_url_len": {
										Type: schema.TypeInt, Optional: true, Default: 128, Description: "Max length of URL log (Max URL length (Default: 128 char))",
									},
									"include_all_headers": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include all configured headers despite of absence in HTTP request",
									},
									"disable_sequence_check": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable http packet sequence check and don't drop out of order packets",
									},
								},
							},
						},
					},
				},
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to the logging template (Service Group Name)",
			},
			"session_periodic_log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval": {
							Type: schema.TypeInt, Optional: true, Description: "Logging time interval (minutes) for long lived sessions",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"severity": {
				Type: schema.TypeString, Optional: true, Default: "informational", Description: "'emergency': 0: Emergency; 'alert': 1: Alert; 'critical': 2: Critical; 'error': 3: Error; 'warning': 4: Warning; 'notice': 5: Notice; 'informational': 6: Informational; 'debug': 7: Debug;",
			},
			"source_address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "Specify source IP address",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify source IPv6 address",
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
func resourceFwTemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFwTemplateLoggingIncludeHttp(d []interface{}) edpt.FwTemplateLoggingIncludeHttp {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingIncludeHttp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HeaderCfg = getSliceFwTemplateLoggingIncludeHttpHeaderCfg(in["header_cfg"].([]interface{}))
		ret.L4SessionInfo = in["l4_session_info"].(int)
		ret.Method = in["method"].(int)
		ret.RequestNumber = in["request_number"].(int)
		ret.FileExtension = in["file_extension"].(int)
	}
	return ret
}

func getSliceFwTemplateLoggingIncludeHttpHeaderCfg(d []interface{}) []edpt.FwTemplateLoggingIncludeHttpHeaderCfg {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingIncludeHttpHeaderCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingIncludeHttpHeaderCfg
		oi.HttpHeader = in["http_header"].(string)
		oi.MaxLength = in["max_length"].(int)
		oi.CustomHeaderName = in["custom_header_name"].(string)
		oi.CustomMaxLength = in["custom_max_length"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwTemplateLoggingIncludeRadiusAttribute(d []interface{}) edpt.FwTemplateLoggingIncludeRadiusAttribute {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingIncludeRadiusAttribute
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AttrCfg = getSliceFwTemplateLoggingIncludeRadiusAttributeAttrCfg(in["attr_cfg"].([]interface{}))
		ret.NoQuote = in["no_quote"].(int)
		ret.FramedIpv6Prefix = in["framed_ipv6_prefix"].(int)
		ret.PrefixLength = in["prefix_length"].(string)
		ret.InsertIfNotExisting = in["insert_if_not_existing"].(int)
		ret.ZeroInCustomAttr = in["zero_in_custom_attr"].(int)
	}
	return ret
}

func getSliceFwTemplateLoggingIncludeRadiusAttributeAttrCfg(d []interface{}) []edpt.FwTemplateLoggingIncludeRadiusAttributeAttrCfg {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingIncludeRadiusAttributeAttrCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingIncludeRadiusAttributeAttrCfg
		oi.Attr = in["attr"].(string)
		oi.AttrEvent = in["attr_event"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwTemplateLoggingLog(d []interface{}) edpt.FwTemplateLoggingLog {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingLog
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpRequests = in["http_requests"].(string)
	}
	return ret
}

func getObjectFwTemplateLoggingRule(d []interface{}) edpt.FwTemplateLoggingRule {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingRule
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleHttpRequests = getObjectFwTemplateLoggingRuleRuleHttpRequests(in["rule_http_requests"].([]interface{}))
	}
	return ret
}

func getObjectFwTemplateLoggingRuleRuleHttpRequests(d []interface{}) edpt.FwTemplateLoggingRuleRuleHttpRequests {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingRuleRuleHttpRequests
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DestPort = getSliceFwTemplateLoggingRuleRuleHttpRequestsDestPort(in["dest_port"].([]interface{}))
		ret.LogEveryHttpRequest = in["log_every_http_request"].(int)
		ret.MaxUrlLen = in["max_url_len"].(int)
		ret.IncludeAllHeaders = in["include_all_headers"].(int)
		ret.DisableSequenceCheck = in["disable_sequence_check"].(int)
	}
	return ret
}

func getSliceFwTemplateLoggingRuleRuleHttpRequestsDestPort(d []interface{}) []edpt.FwTemplateLoggingRuleRuleHttpRequestsDestPort {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingRuleRuleHttpRequestsDestPort, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingRuleRuleHttpRequestsDestPort
		oi.DestPortNumber = in["dest_port_number"].(int)
		oi.IncludeByteCount = in["include_byte_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwTemplateLoggingSessionPeriodicLog378(d []interface{}) edpt.FwTemplateLoggingSessionPeriodicLog378 {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingSessionPeriodicLog378
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		//omit uuid
	}
	return ret
}

func getObjectFwTemplateLoggingSourceAddress379(d []interface{}) edpt.FwTemplateLoggingSourceAddress379 {

	count1 := len(d)
	var ret edpt.FwTemplateLoggingSourceAddress379
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointFwTemplateLogging(d *schema.ResourceData) edpt.FwTemplateLogging {
	var ret edpt.FwTemplateLogging
	ret.Inst.Facility = d.Get("facility").(string)
	ret.Inst.Format = d.Get("format").(string)
	ret.Inst.IncludeDestFqdn = d.Get("include_dest_fqdn").(int)
	ret.Inst.IncludeHttp = getObjectFwTemplateLoggingIncludeHttp(d.Get("include_http").([]interface{}))
	ret.Inst.IncludeRadiusAttribute = getObjectFwTemplateLoggingIncludeRadiusAttribute(d.Get("include_radius_attribute").([]interface{}))
	ret.Inst.Log = getObjectFwTemplateLoggingLog(d.Get("log").([]interface{}))
	ret.Inst.MergedStyle = d.Get("merged_style").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Resolution = d.Get("resolution").(string)
	ret.Inst.Rule = getObjectFwTemplateLoggingRule(d.Get("rule").([]interface{}))
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.SessionPeriodicLog = getObjectFwTemplateLoggingSessionPeriodicLog378(d.Get("session_periodic_log").([]interface{}))
	ret.Inst.Severity = d.Get("severity").(string)
	ret.Inst.SourceAddress = getObjectFwTemplateLoggingSourceAddress379(d.Get("source_address").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
