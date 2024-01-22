package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_http`: HTTP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateHttpCreate,
		UpdateContext: resourceSlbTemplateHttpUpdate,
		ReadContext:   resourceSlbTemplateHttpRead,
		DeleteContext: resourceSlbTemplateHttpDelete,

		Schema: map[string]*schema.Schema{
			"allowed_methods": {
				Type: schema.TypeString, Optional: true, Description: "Enable allowed-method check (List of allowed HTTP methods)",
			},
			"allowed_methods_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Respond 400 directly;",
			},
			"bypass_sg": {
				Type: schema.TypeString, Optional: true, Description: "Select service group for non-http traffic (Service Group Name)",
			},
			"client_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client session timeout if the next request is not received (timeout in seconds. 0 means disable, default is 0)",
			},
			"client_ip_hdr_replace": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace the existing header",
			},
			"client_port_hdr_replace": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace the existing header",
			},
			"compression_auto_disable_on_high_cpu": {
				Type: schema.TypeInt, Optional: true, Description: "Auto-disable software compression on high cpu usage (Disable compression if cpu usage is above threshold. Default is off.)",
			},
			"compression_br_level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "brotli compression level, default 1 (brotli compression level value, default is 1)",
			},
			"compression_br_sliding_window_size": {
				Type: schema.TypeInt, Optional: true, Description: "brotli compression sliding window size, default 10 (brotli compression sliding window size in the form of log (i.e., 10 means 1k-16MB bytes))",
			},
			"compression_content_type": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_type": {
							Type: schema.TypeString, Optional: true, Description: "Compression content-type",
						},
					},
				},
			},
			"compression_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Compression",
			},
			"compression_exclude_content_type": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_content_type": {
							Type: schema.TypeString, Optional: true, Description: "Compression exclude content-type (Compression exclude content type)",
						},
					},
				},
			},
			"compression_exclude_uri": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_uri": {
							Type: schema.TypeString, Optional: true, Description: "Compression exclude uri",
						},
					},
				},
			},
			"compression_keep_accept_encoding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep accept encoding",
			},
			"compression_keep_accept_encoding_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Server Accept Encoding",
			},
			"compression_level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "gzip compression level, default 1 (gzip compression level value, default is 1)",
			},
			"compression_method_order": {
				Type: schema.TypeString, Optional: true, Description: "Method Order (Order to decide which compression algorithm to be applied when multiple algorithms are acceptable)",
			},
			"compression_minimum_content_length": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Minimum Content Length (Minimum content length for compression in bytes. Default is 120.)",
			},
			"cont_wait_for_req_complete100": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "When REQ has Expect 100 and response is not 100, then wait for whole request to be sent",
			},
			"cookie_format": {
				Type: schema.TypeString, Optional: true, Description: "'rfc6265': Follow rfc6265;",
			},
			"cookie_samesite": {
				Type: schema.TypeString, Optional: true, Description: "'none': none; 'lax': lax; 'strict': strict;",
			},
			"default_charset": {
				Type: schema.TypeString, Optional: true, Default: "utf-8", Description: "'iso-8859-1': Use ISO-8859-1 as the default charset; 'utf-8': Use UTF-8 as the default charset; 'us-ascii': Use US-ASCII as the default charset;",
			},
			"disallowed_methods": {
				Type: schema.TypeString, Optional: true, Description: "Enable disallowed-method check (List of disallowed HTTP methods)",
			},
			"disallowed_methods_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Respond 400 directly;",
			},
			"failover_url": {
				Type: schema.TypeString, Optional: true, Description: "Failover to this URL (Failover URL Name)",
			},
			"frame_limit": {
				Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Limit the number of CONTINUATION, PING, PRIORITY, RESET, SETTINGS and empty frames in one HTTP2 connection, default 10000",
			},
			"host_switching": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_switching_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Select service group if hostname contains another string; 'ends-with': Select service group if hostname ends with another string; 'equals': Select service group if hostname equals another string; 'starts-with': Select service group if hostname starts with another string; 'regex-match': Select service group if URL string matches with regular expression; 'host-hits-enable': Enables Host Hits counters;",
						},
						"host_match_string": {
							Type: schema.TypeString, Optional: true, Description: "Hostname String",
						},
						"host_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Create a Service Group comprising Servers (Service Group Name)",
						},
					},
				},
			},
			"http_protocol_check": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"h2up_content_length_alias": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"malformed_h2up_header_value": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"malformed_h2up_scheme_value": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"h2up_with_transfer_encoding": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"multiple_content_length": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"multiple_transfer_encoding": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"transfer_encoding_and_content_length": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and Send 400 to the client side;",
						},
						"get_and_payload": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"h2up_with_host_and_auth": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"header_filter_rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": {
										Type: schema.TypeInt, Required: true, Description: "Specify a sequence number",
									},
									"match_type_value": {
										Type: schema.TypeString, Optional: true, Description: "'full-text': Full text match; 'pcre': PCRE match;",
									},
									"header_name_value": {
										Type: schema.TypeString, Optional: true, Description: "Header name value",
									},
									"header_value_value": {
										Type: schema.TypeString, Optional: true, Description: "Header value",
									},
									"action_value": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request;",
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
			"http2_client_no_snat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set max-concurrent-stream = 1 when the client side is HTTP2 and no source-nat configuration is under vport",
			},
			"insert_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert Client IP address into HTTP header",
			},
			"insert_client_ip_header_name": {
				Type: schema.TypeString, Optional: true, Description: "HTTP Header Name for inserting Client IP",
			},
			"insert_client_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert Client Port address into HTTP header",
			},
			"insert_client_port_header_name": {
				Type: schema.TypeString, Optional: true, Description: "HTTP Header Name for inserting Client Port",
			},
			"keep_client_alive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep client alive",
			},
			"log_retry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "log when HTTP request retry",
			},
			"max_concurrent_streams": {
				Type: schema.TypeInt, Optional: true, Default: 50, Description: "(http2 only) Max concurrent streams, default 50",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "HTTP Template Name",
			},
			"non_http_bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bypass non-http traffic instead of dropping",
			},
			"persist_on_401": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persist to the same server if the response code is 401",
			},
			"prefix": {
				Type: schema.TypeString, Optional: true, Description: "'host': the cookie will have been set with a Secure attribute, a Path attribute with a value of /, and no Domain attribute; 'secure': the cookie will have been set with a Secure attribute; 'check': check server prefix and enforce prefix format;",
			},
			"rd_port": {
				Type: schema.TypeInt, Optional: true, Description: "Port (Port Number)",
			},
			"rd_resp_code": {
				Type: schema.TypeString, Optional: true, Description: "'301': Moved Permanently; '302': Found; '303': See Other; '307': Temporary Redirect;",
			},
			"rd_secure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use HTTPS",
			},
			"rd_simple_loc": {
				Type: schema.TypeString, Optional: true, Description: "Redirect location tag absolute URI string",
			},
			"redirect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Automatically send a redirect response",
			},
			"redirect_rewrite": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redirect_match": {
										Type: schema.TypeString, Optional: true, Description: "URL Matching (Pattern URL String)",
									},
									"rewrite_to": {
										Type: schema.TypeString, Optional: true, Description: "Rewrite to Destination URL String",
									},
								},
							},
						},
						"redirect_secure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use HTTPS",
						},
						"redirect_secure_port": {
							Type: schema.TypeInt, Optional: true, Default: 443, Description: "Port (Port Number)",
						},
					},
				},
			},
			"req_hdr_wait_time": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP request header wait time before abort connection",
			},
			"req_hdr_wait_time_val": {
				Type: schema.TypeInt, Optional: true, Default: 7, Description: "Number of seconds wait for client request header (default is 7)",
			},
			"request_header_erase_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_header_erase": {
							Type: schema.TypeString, Optional: true, Description: "Erase a header from HTTP request (Header Name)",
						},
					},
				},
			},
			"request_header_insert_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_header_insert": {
							Type: schema.TypeString, Optional: true, Description: "Insert a header into HTTP request (Header Content (Format: \"[name]:[value]\"))",
						},
						"request_header_insert_type": {
							Type: schema.TypeString, Optional: true, Description: "'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;",
						},
					},
				},
			},
			"request_line_case_insensitive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Parse http request line as case insensitive",
			},
			"request_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Request timeout if response not received (timeout in seconds)",
			},
			"response_content_replace_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_content_replace": {
							Type: schema.TypeString, Optional: true, Description: "replace the data from HTTP response content (String in the http content need to be replaced)",
						},
						"response_new_string": {
							Type: schema.TypeString, Optional: true, Description: "String will be in the http content",
						},
					},
				},
			},
			"response_header_erase_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_header_erase": {
							Type: schema.TypeString, Optional: true, Description: "Erase a header from HTTP response (Header Name)",
						},
					},
				},
			},
			"response_header_insert_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_header_insert": {
							Type: schema.TypeString, Optional: true, Description: "Insert a header into HTTP response (Header Content (Format: \"[name]:[value]\"))",
						},
						"response_header_insert_type": {
							Type: schema.TypeString, Optional: true, Description: "'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;",
						},
					},
				},
			},
			"retry_on_5xx": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Retry http request on HTTP 5xx code and request timeout",
			},
			"retry_on_5xx_per_req": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Retry http request on HTTP 5xx code for each request",
			},
			"retry_on_5xx_per_req_val": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of times to retry (default is 3)",
			},
			"retry_on_5xx_val": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Number of times to retry (default is 3)",
			},
			"server_support_http2_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Notify the vport regarding whether server supports http2 only",
			},
			"server_support_http2_only_value": {
				Type: schema.TypeString, Optional: true, Default: "auto-detect", Description: "'auto-detect': Commuincate with the server via HTTP/2 when an support-http2-only rport is detected; 'force': Communicate with the server via HTTP/2 when possible;",
			},
			"stream_cancellation_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "cancellation limit, default 0 (accumulated cancellation limit value, default is 0)",
			},
			"stream_cancellation_rate": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "cancellation rate, default 10 (cancellation rate value, default is 10)",
			},
			"strict_transaction_switch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force server selection on every HTTP request",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "Logging template (Logging Config name)",
						},
					},
				},
			},
			"term_11client_hdr_conn_close": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Terminate HTTP 1.1 client when req has Connection: close",
			},
			"url_hash_first": {
				Type: schema.TypeInt, Optional: true, Description: "Use the begining part of URL to calculate hash value (URL string length to calculate hash value)",
			},
			"url_hash_last": {
				Type: schema.TypeInt, Optional: true, Description: "Use the end part of URL to calculate hash value (URL string length to calculate hash value)",
			},
			"url_hash_offset": {
				Type: schema.TypeInt, Optional: true, Description: "Skip part of URL to calculate hash value (Offset of the URL string)",
			},
			"url_hash_persist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use URL's hash value to select server",
			},
			"url_switching": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url_switching_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': Select service group if URL string contains another string; 'ends-with': Select service group if URL string ends with another string; 'equals': Select service group if URL string equals another string; 'starts-with': Select service group if URL string starts with another string; 'regex-match': Select service group if URL string matches with regular expression; 'url-case-insensitive': Case insensitive URL switching; 'url-hits-enable': Enables URL Hits;",
						},
						"url_match_string": {
							Type: schema.TypeString, Optional: true, Description: "URL String",
						},
						"url_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Create a Service Group comprising Servers (Service Group Name)",
						},
					},
				},
			},
			"use_server_status": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Server-Status header to do URL hashing",
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
func resourceSlbTemplateHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateHttpCompressionContentType(d []interface{}) []edpt.SlbTemplateHttpCompressionContentType {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpCompressionContentType, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpCompressionContentType
		oi.ContentType = in["content_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpCompressionExcludeContentType(d []interface{}) []edpt.SlbTemplateHttpCompressionExcludeContentType {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpCompressionExcludeContentType, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpCompressionExcludeContentType
		oi.ExcludeContentType = in["exclude_content_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpCompressionExcludeUri(d []interface{}) []edpt.SlbTemplateHttpCompressionExcludeUri {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpCompressionExcludeUri, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpCompressionExcludeUri
		oi.ExcludeUri = in["exclude_uri"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpHostSwitching(d []interface{}) []edpt.SlbTemplateHttpHostSwitching {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpHostSwitching, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpHostSwitching
		oi.HostSwitchingType = in["host_switching_type"].(string)
		oi.HostMatchString = in["host_match_string"].(string)
		oi.HostServiceGroup = in["host_service_group"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateHttpHttpProtocolCheck1447(d []interface{}) edpt.SlbTemplateHttpHttpProtocolCheck1447 {

	count1 := len(d)
	var ret edpt.SlbTemplateHttpHttpProtocolCheck1447
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.H2upContentLengthAlias = in["h2up_content_length_alias"].(string)
		ret.MalformedH2upHeaderValue = in["malformed_h2up_header_value"].(string)
		ret.MalformedH2upSchemeValue = in["malformed_h2up_scheme_value"].(string)
		ret.H2upWithTransferEncoding = in["h2up_with_transfer_encoding"].(string)
		ret.MultipleContentLength = in["multiple_content_length"].(string)
		ret.MultipleTransferEncoding = in["multiple_transfer_encoding"].(string)
		ret.TransferEncodingAndContentLength = in["transfer_encoding_and_content_length"].(string)
		ret.GetAndPayload = in["get_and_payload"].(string)
		ret.H2upWithHostAndAuth = in["h2up_with_host_and_auth"].(string)
		//omit uuid
		ret.HeaderFilterRuleList = getSliceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448(in["header_filter_rule_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448(d []interface{}) []edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448
		oi.SeqNum = in["seq_num"].(int)
		oi.MatchTypeValue = in["match_type_value"].(string)
		oi.HeaderNameValue = in["header_name_value"].(string)
		oi.HeaderValueValue = in["header_value_value"].(string)
		oi.ActionValue = in["action_value"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateHttpRedirectRewrite(d []interface{}) edpt.SlbTemplateHttpRedirectRewrite {

	count1 := len(d)
	var ret edpt.SlbTemplateHttpRedirectRewrite
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MatchList = getSliceSlbTemplateHttpRedirectRewriteMatchList(in["match_list"].([]interface{}))
		ret.RedirectSecure = in["redirect_secure"].(int)
		ret.RedirectSecurePort = in["redirect_secure_port"].(int)
	}
	return ret
}

func getSliceSlbTemplateHttpRedirectRewriteMatchList(d []interface{}) []edpt.SlbTemplateHttpRedirectRewriteMatchList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpRedirectRewriteMatchList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpRedirectRewriteMatchList
		oi.RedirectMatch = in["redirect_match"].(string)
		oi.RewriteTo = in["rewrite_to"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpRequestHeaderEraseList(d []interface{}) []edpt.SlbTemplateHttpRequestHeaderEraseList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpRequestHeaderEraseList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpRequestHeaderEraseList
		oi.RequestHeaderErase = in["request_header_erase"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpRequestHeaderInsertList(d []interface{}) []edpt.SlbTemplateHttpRequestHeaderInsertList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpRequestHeaderInsertList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpRequestHeaderInsertList
		oi.RequestHeaderInsert = in["request_header_insert"].(string)
		oi.RequestHeaderInsertType = in["request_header_insert_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpResponseContentReplaceList(d []interface{}) []edpt.SlbTemplateHttpResponseContentReplaceList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpResponseContentReplaceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpResponseContentReplaceList
		oi.ResponseContentReplace = in["response_content_replace"].(string)
		oi.ResponseNewString = in["response_new_string"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpResponseHeaderEraseList(d []interface{}) []edpt.SlbTemplateHttpResponseHeaderEraseList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpResponseHeaderEraseList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpResponseHeaderEraseList
		oi.ResponseHeaderErase = in["response_header_erase"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateHttpResponseHeaderInsertList(d []interface{}) []edpt.SlbTemplateHttpResponseHeaderInsertList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpResponseHeaderInsertList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpResponseHeaderInsertList
		oi.ResponseHeaderInsert = in["response_header_insert"].(string)
		oi.ResponseHeaderInsertType = in["response_header_insert_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateHttpTemplate(d []interface{}) edpt.SlbTemplateHttpTemplate {

	count1 := len(d)
	var ret edpt.SlbTemplateHttpTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceSlbTemplateHttpUrlSwitching(d []interface{}) []edpt.SlbTemplateHttpUrlSwitching {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpUrlSwitching, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpUrlSwitching
		oi.UrlSwitchingType = in["url_switching_type"].(string)
		oi.UrlMatchString = in["url_match_string"].(string)
		oi.UrlServiceGroup = in["url_service_group"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateHttp(d *schema.ResourceData) edpt.SlbTemplateHttp {
	var ret edpt.SlbTemplateHttp
	ret.Inst.AllowedMethods = d.Get("allowed_methods").(string)
	ret.Inst.AllowedMethodsAction = d.Get("allowed_methods_action").(string)
	ret.Inst.BypassSg = d.Get("bypass_sg").(string)
	ret.Inst.ClientIdleTimeout = d.Get("client_idle_timeout").(int)
	ret.Inst.ClientIpHdrReplace = d.Get("client_ip_hdr_replace").(int)
	ret.Inst.ClientPortHdrReplace = d.Get("client_port_hdr_replace").(int)
	ret.Inst.CompressionAutoDisableOnHighCpu = d.Get("compression_auto_disable_on_high_cpu").(int)
	ret.Inst.CompressionBrLevel = d.Get("compression_br_level").(int)
	ret.Inst.CompressionBrSlidingWindowSize = d.Get("compression_br_sliding_window_size").(int)
	ret.Inst.CompressionContentType = getSliceSlbTemplateHttpCompressionContentType(d.Get("compression_content_type").([]interface{}))
	ret.Inst.CompressionEnable = d.Get("compression_enable").(int)
	ret.Inst.CompressionExcludeContentType = getSliceSlbTemplateHttpCompressionExcludeContentType(d.Get("compression_exclude_content_type").([]interface{}))
	ret.Inst.CompressionExcludeUri = getSliceSlbTemplateHttpCompressionExcludeUri(d.Get("compression_exclude_uri").([]interface{}))
	ret.Inst.CompressionKeepAcceptEncoding = d.Get("compression_keep_accept_encoding").(int)
	ret.Inst.CompressionKeepAcceptEncodingEnable = d.Get("compression_keep_accept_encoding_enable").(int)
	ret.Inst.CompressionLevel = d.Get("compression_level").(int)
	ret.Inst.CompressionMethodOrder = d.Get("compression_method_order").(string)
	ret.Inst.CompressionMinimumContentLength = d.Get("compression_minimum_content_length").(int)
	ret.Inst.ContWaitForReqComplete100 = d.Get("cont_wait_for_req_complete100").(int)
	ret.Inst.CookieFormat = d.Get("cookie_format").(string)
	ret.Inst.CookieSamesite = d.Get("cookie_samesite").(string)
	ret.Inst.DefaultCharset = d.Get("default_charset").(string)
	ret.Inst.DisallowedMethods = d.Get("disallowed_methods").(string)
	ret.Inst.DisallowedMethodsAction = d.Get("disallowed_methods_action").(string)
	ret.Inst.FailoverUrl = d.Get("failover_url").(string)
	ret.Inst.FrameLimit = d.Get("frame_limit").(int)
	ret.Inst.HostSwitching = getSliceSlbTemplateHttpHostSwitching(d.Get("host_switching").([]interface{}))
	ret.Inst.HttpProtocolCheck = getObjectSlbTemplateHttpHttpProtocolCheck1447(d.Get("http_protocol_check").([]interface{}))
	ret.Inst.Http2ClientNoSnat = d.Get("http2_client_no_snat").(int)
	ret.Inst.InsertClientIp = d.Get("insert_client_ip").(int)
	ret.Inst.InsertClientIpHeaderName = d.Get("insert_client_ip_header_name").(string)
	ret.Inst.InsertClientPort = d.Get("insert_client_port").(int)
	ret.Inst.InsertClientPortHeaderName = d.Get("insert_client_port_header_name").(string)
	ret.Inst.KeepClientAlive = d.Get("keep_client_alive").(int)
	ret.Inst.LogRetry = d.Get("log_retry").(int)
	ret.Inst.MaxConcurrentStreams = d.Get("max_concurrent_streams").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NonHttpBypass = d.Get("non_http_bypass").(int)
	ret.Inst.PersistOn401 = d.Get("persist_on_401").(int)
	ret.Inst.Prefix = d.Get("prefix").(string)
	ret.Inst.RdPort = d.Get("rd_port").(int)
	ret.Inst.RdRespCode = d.Get("rd_resp_code").(string)
	ret.Inst.RdSecure = d.Get("rd_secure").(int)
	ret.Inst.RdSimpleLoc = d.Get("rd_simple_loc").(string)
	ret.Inst.Redirect = d.Get("redirect").(int)
	ret.Inst.RedirectRewrite = getObjectSlbTemplateHttpRedirectRewrite(d.Get("redirect_rewrite").([]interface{}))
	ret.Inst.ReqHdrWaitTime = d.Get("req_hdr_wait_time").(int)
	ret.Inst.ReqHdrWaitTimeVal = d.Get("req_hdr_wait_time_val").(int)
	ret.Inst.RequestHeaderEraseList = getSliceSlbTemplateHttpRequestHeaderEraseList(d.Get("request_header_erase_list").([]interface{}))
	ret.Inst.RequestHeaderInsertList = getSliceSlbTemplateHttpRequestHeaderInsertList(d.Get("request_header_insert_list").([]interface{}))
	ret.Inst.RequestLineCaseInsensitive = d.Get("request_line_case_insensitive").(int)
	ret.Inst.RequestTimeout = d.Get("request_timeout").(int)
	ret.Inst.ResponseContentReplaceList = getSliceSlbTemplateHttpResponseContentReplaceList(d.Get("response_content_replace_list").([]interface{}))
	ret.Inst.ResponseHeaderEraseList = getSliceSlbTemplateHttpResponseHeaderEraseList(d.Get("response_header_erase_list").([]interface{}))
	ret.Inst.ResponseHeaderInsertList = getSliceSlbTemplateHttpResponseHeaderInsertList(d.Get("response_header_insert_list").([]interface{}))
	ret.Inst.RetryOn5xx = d.Get("retry_on_5xx").(int)
	ret.Inst.RetryOn5xxPerReq = d.Get("retry_on_5xx_per_req").(int)
	ret.Inst.RetryOn5xxPerReqVal = d.Get("retry_on_5xx_per_req_val").(int)
	ret.Inst.RetryOn5xxVal = d.Get("retry_on_5xx_val").(int)
	ret.Inst.ServerSupportHttp2Only = d.Get("server_support_http2_only").(int)
	ret.Inst.ServerSupportHttp2OnlyValue = d.Get("server_support_http2_only_value").(string)
	ret.Inst.StreamCancellationLimit = d.Get("stream_cancellation_limit").(int)
	ret.Inst.StreamCancellationRate = d.Get("stream_cancellation_rate").(int)
	ret.Inst.StrictTransactionSwitch = d.Get("strict_transaction_switch").(int)
	ret.Inst.Template = getObjectSlbTemplateHttpTemplate(d.Get("template").([]interface{}))
	ret.Inst.Term11clientHdrConnClose = d.Get("term_11client_hdr_conn_close").(int)
	ret.Inst.UrlHashFirst = d.Get("url_hash_first").(int)
	ret.Inst.UrlHashLast = d.Get("url_hash_last").(int)
	ret.Inst.UrlHashOffset = d.Get("url_hash_offset").(int)
	ret.Inst.UrlHashPersist = d.Get("url_hash_persist").(int)
	ret.Inst.UrlSwitching = getSliceSlbTemplateHttpUrlSwitching(d.Get("url_switching").([]interface{}))
	ret.Inst.UseServerStatus = d.Get("use_server_status").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
