package thunder

//Thunder resource SlbTemplateHttp

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateHttp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateHttpCreate,
		UpdateContext: resourceSlbTemplateHttpUpdate,
		ReadContext:   resourceSlbTemplateHttpRead,
		DeleteContext: resourceSlbTemplateHttpDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"compression_auto_disable_on_high_cpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compression_content_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"compression_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compression_exclude_content_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_content_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"compression_exclude_uri": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclude_uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"compression_keep_accept_encoding": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compression_keep_accept_encoding_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compression_level": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"compression_minimum_content_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"default_charset": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"max_concurrent_streams": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"frame_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"failover_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_switching_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host_service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_ip_header_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_ip_hdr_replace": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_port_header_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_port_hdr_replace": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"non_http_bypass": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bypass_sg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rd_simple_loc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rd_secure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rd_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rd_resp_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"redirect_rewrite": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redirect_match": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"rewrite_to": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"redirect_secure": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"redirect_secure_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"request_header_erase_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_header_erase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"request_header_insert_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_header_insert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"request_header_insert_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"response_content_replace_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_content_replace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"response_new_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"response_header_erase_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_header_erase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"response_header_insert_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_header_insert": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"response_header_insert_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry_on_5xx": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry_on_5xx_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry_on_5xx_per_req": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry_on_5xx_per_req_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"strict_transaction_switch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"term_11client_hdr_conn_close": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persist_on_401": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"100_cont_wait_for_req_complete": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_persist": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_first": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_last": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_server_status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url_switching_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"req_hdr_wait_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"req_hdr_wait_time_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"request_line_case_insensitive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"keep_client_alive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cookie_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cookie_samesite": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateHttp (Inside resourceSlbTemplateHttpCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateHttp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateHttp --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateHttp(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateHttpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateHttp (Inside resourceSlbTemplateHttpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateHttp(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceSlbTemplateHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateHttp   (Inside resourceSlbTemplateHttpUpdate) ")
		data := dataToSlbTemplateHttp(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateHttp ")
		err := go_thunder.PutSlbTemplateHttp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateHttpRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateHttpDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateHttp(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateHttp(d *schema.ResourceData) go_thunder.SlbTemplateHttp {
	var vc go_thunder.SlbTemplateHttp
	var c go_thunder.SlbTemplateHTTPInstance
	c.SlbTemplateHTTPInstanceName = d.Get("name").(string)
	c.SlbTemplateHTTPInstanceCompressionAutoDisableOnHighCPU = d.Get("compression_auto_disable_on_high_cpu").(int)

	SlbTemplateHTTPInstanceCompressionContentTypeCount := d.Get("compression_content_type.#").(int)
	c.SlbTemplateHTTPInstanceCompressionContentTypeContentType = make([]go_thunder.SlbTemplateHTTPInstanceCompressionContentType, 0, SlbTemplateHTTPInstanceCompressionContentTypeCount)

	for i := 0; i < SlbTemplateHTTPInstanceCompressionContentTypeCount; i++ {
		var obj1 go_thunder.SlbTemplateHTTPInstanceCompressionContentType
		prefix1 := fmt.Sprintf("compression_content_type.%d.", i)
		obj1.SlbTemplateHTTPInstanceCompressionContentTypeContentType = d.Get(prefix1 + "content_type").(string)
		c.SlbTemplateHTTPInstanceCompressionContentTypeContentType = append(c.SlbTemplateHTTPInstanceCompressionContentTypeContentType, obj1)
	}

	c.SlbTemplateHTTPInstanceCompressionEnable = d.Get("compression_enable").(int)

	SlbTemplateHTTPInstanceCompressionExcludeContentTypeCount := d.Get("compression_exclude_content_type.#").(int)
	c.SlbTemplateHTTPInstanceCompressionExcludeContentTypeExcludeContentType = make([]go_thunder.SlbTemplateHTTPInstanceCompressionExcludeContentType, 0, SlbTemplateHTTPInstanceCompressionExcludeContentTypeCount)

	for i := 0; i < SlbTemplateHTTPInstanceCompressionExcludeContentTypeCount; i++ {
		var obj2 go_thunder.SlbTemplateHTTPInstanceCompressionExcludeContentType
		prefix2 := fmt.Sprintf("compression_exclude_content_type.%d.", i)
		obj2.SlbTemplateHTTPInstanceCompressionExcludeContentTypeExcludeContentType = d.Get(prefix2 + "exclude_content_type").(string)
		c.SlbTemplateHTTPInstanceCompressionExcludeContentTypeExcludeContentType = append(c.SlbTemplateHTTPInstanceCompressionExcludeContentTypeExcludeContentType, obj2)
	}

	SlbTemplateHTTPInstanceCompressionExcludeURICount := d.Get("compression_exclude_uri.#").(int)
	c.SlbTemplateHTTPInstanceCompressionExcludeURIExcludeURI = make([]go_thunder.SlbTemplateHTTPInstanceCompressionExcludeURI, 0, SlbTemplateHTTPInstanceCompressionExcludeURICount)

	for i := 0; i < SlbTemplateHTTPInstanceCompressionExcludeURICount; i++ {
		var obj3 go_thunder.SlbTemplateHTTPInstanceCompressionExcludeURI
		prefix3 := fmt.Sprintf("compression_exclude_uri.%d.", i)
		obj3.SlbTemplateHTTPInstanceCompressionExcludeURIExcludeURI = d.Get(prefix3 + "exclude_uri").(string)
		c.SlbTemplateHTTPInstanceCompressionExcludeURIExcludeURI = append(c.SlbTemplateHTTPInstanceCompressionExcludeURIExcludeURI, obj3)
	}

	c.SlbTemplateHTTPInstanceCompressionKeepAcceptEncoding = d.Get("compression_keep_accept_encoding").(int)
	c.SlbTemplateHTTPInstanceCompressionKeepAcceptEncodingEnable = d.Get("compression_keep_accept_encoding_enable").(int)
	c.SlbTemplateHTTPInstanceCompressionLevel = d.Get("compression_level").(int)
	c.SlbTemplateHTTPInstanceCompressionMinimumContentLength = d.Get("compression_minimum_content_length").(int)
	c.SlbTemplateHTTPInstanceDefaultCharset = d.Get("default_charset").(string)
	c.SlbTemplateHTTPInstanceMaxConcurrentStreams = d.Get("max_concurrent_streams").(int)
	c.SlbTemplateHTTPInstanceFrameLimit = d.Get("frame_limit").(int)
	c.SlbTemplateHTTPInstanceFailoverUrl = d.Get("failover_url").(string)

	SlbTemplateHTTPInstanceHostSwitchingCount := d.Get("host_switching.#").(int)
	c.SlbTemplateHTTPInstanceHostSwitchingHostSwitchingType = make([]go_thunder.SlbTemplateHTTPInstanceHostSwitching, 0, SlbTemplateHTTPInstanceHostSwitchingCount)

	for i := 0; i < SlbTemplateHTTPInstanceHostSwitchingCount; i++ {
		var obj4 go_thunder.SlbTemplateHTTPInstanceHostSwitching
		prefix4 := fmt.Sprintf("host_switching.%d.", i)
		obj4.SlbTemplateHTTPInstanceHostSwitchingHostSwitchingType = d.Get(prefix4 + "host_switching_type").(string)
		obj4.SlbTemplateHTTPInstanceHostSwitchingHostMatchString = d.Get(prefix4 + "host_match_string").(string)
		obj4.SlbTemplateHTTPInstanceHostSwitchingHostServiceGroup = d.Get(prefix4 + "host_service_group").(string)
		c.SlbTemplateHTTPInstanceHostSwitchingHostSwitchingType = append(c.SlbTemplateHTTPInstanceHostSwitchingHostSwitchingType, obj4)
	}

	c.SlbTemplateHTTPInstanceInsertClientIP = d.Get("insert_client_ip").(int)
	c.SlbTemplateHTTPInstanceInsertClientIPHeaderName = d.Get("insert_client_ip_header_name").(string)
	c.SlbTemplateHTTPInstanceClientIPHdrReplace = d.Get("client_ip_hdr_replace").(int)
	c.SlbTemplateHTTPInstanceInsertClientPort = d.Get("insert_client_port").(int)
	c.SlbTemplateHTTPInstanceInsertClientPortHeaderName = d.Get("insert_client_port_header_name").(string)
	c.SlbTemplateHTTPInstanceClientPortHdrReplace = d.Get("client_port_hdr_replace").(int)
	c.SlbTemplateHTTPInstanceLogRetry = d.Get("log_retry").(int)
	c.SlbTemplateHTTPInstanceNonHTTPBypass = d.Get("non_http_bypass").(int)
	c.SlbTemplateHTTPInstanceBypassSg = d.Get("bypass_sg").(string)
	c.SlbTemplateHTTPInstanceRedirect = d.Get("redirect").(int)
	c.SlbTemplateHTTPInstanceRdSimpleLoc = d.Get("rd_simple_loc").(string)
	c.SlbTemplateHTTPInstanceRdSecure = d.Get("rd_secure").(int)
	c.SlbTemplateHTTPInstanceRdPort = d.Get("rd_port").(int)
	c.SlbTemplateHTTPInstanceRdRespCode = d.Get("rd_resp_code").(string)

	var obj5 go_thunder.SlbTemplateHTTPInstanceRedirectRewrite
	prefix5 := "redirect_rewrite.0."

	SlbTemplateHTTPInstanceRedirectRewriteMatchListCount := d.Get(prefix5 + "match_list.#").(int)
	obj5.SlbTemplateHTTPInstanceRedirectRewriteMatchListRedirectMatch = make([]go_thunder.SlbTemplateHTTPInstanceRedirectRewriteMatchList, 0, SlbTemplateHTTPInstanceRedirectRewriteMatchListCount)

	for i := 0; i < SlbTemplateHTTPInstanceRedirectRewriteMatchListCount; i++ {
		var obj5_1 go_thunder.SlbTemplateHTTPInstanceRedirectRewriteMatchList
		prefix5_1 := prefix5 + fmt.Sprintf("match_list.%d.", i)
		obj5_1.SlbTemplateHTTPInstanceRedirectRewriteMatchListRedirectMatch = d.Get(prefix5_1 + "redirect_match").(string)
		obj5_1.SlbTemplateHTTPInstanceRedirectRewriteMatchListRewriteTo = d.Get(prefix5_1 + "rewrite_to").(string)
		obj5.SlbTemplateHTTPInstanceRedirectRewriteMatchListRedirectMatch = append(obj5.SlbTemplateHTTPInstanceRedirectRewriteMatchListRedirectMatch, obj5_1)
	}

	obj5.SlbTemplateHTTPInstanceRedirectRewriteRedirectSecure = d.Get(prefix5 + "redirect_secure").(int)
	obj5.SlbTemplateHTTPInstanceRedirectRewriteRedirectSecurePort = d.Get(prefix5 + "redirect_secure_port").(int)

	c.SlbTemplateHTTPInstanceRedirectRewriteMatchList = obj5

	SlbTemplateHTTPInstanceRequestHeaderEraseListCount := d.Get("request_header_erase_list.#").(int)
	c.SlbTemplateHTTPInstanceRequestHeaderEraseListRequestHeaderErase = make([]go_thunder.SlbTemplateHTTPInstanceRequestHeaderEraseList, 0, SlbTemplateHTTPInstanceRequestHeaderEraseListCount)

	for i := 0; i < SlbTemplateHTTPInstanceRequestHeaderEraseListCount; i++ {
		var obj6 go_thunder.SlbTemplateHTTPInstanceRequestHeaderEraseList
		prefix6 := fmt.Sprintf("request_header_erase_list.%d.", i)
		obj6.SlbTemplateHTTPInstanceRequestHeaderEraseListRequestHeaderErase = d.Get(prefix6 + "request_header_erase").(string)
		c.SlbTemplateHTTPInstanceRequestHeaderEraseListRequestHeaderErase = append(c.SlbTemplateHTTPInstanceRequestHeaderEraseListRequestHeaderErase, obj6)
	}

	SlbTemplateHTTPInstanceRequestHeaderInsertListCount := d.Get("request_header_insert_list.#").(int)
	c.SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsert = make([]go_thunder.SlbTemplateHTTPInstanceRequestHeaderInsertList, 0, SlbTemplateHTTPInstanceRequestHeaderInsertListCount)

	for i := 0; i < SlbTemplateHTTPInstanceRequestHeaderInsertListCount; i++ {
		var obj7 go_thunder.SlbTemplateHTTPInstanceRequestHeaderInsertList
		prefix7 := fmt.Sprintf("request_header_insert_list.%d.", i)
		obj7.SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsert = d.Get(prefix7 + "request_header_insert").(string)
		obj7.SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsertType = d.Get(prefix7 + "request_header_insert_type").(string)
		c.SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsert = append(c.SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsert, obj7)
	}

	SlbTemplateHTTPInstanceResponseContentReplaceListCount := d.Get("response_content_replace_list.#").(int)
	c.SlbTemplateHTTPInstanceResponseContentReplaceListResponseContentReplace = make([]go_thunder.SlbTemplateHTTPInstanceResponseContentReplaceList, 0, SlbTemplateHTTPInstanceResponseContentReplaceListCount)

	for i := 0; i < SlbTemplateHTTPInstanceResponseContentReplaceListCount; i++ {
		var obj8 go_thunder.SlbTemplateHTTPInstanceResponseContentReplaceList
		prefix8 := fmt.Sprintf("response_content_replace_list.%d.", i)
		obj8.SlbTemplateHTTPInstanceResponseContentReplaceListResponseContentReplace = d.Get(prefix8 + "response_content_replace").(string)
		obj8.SlbTemplateHTTPInstanceResponseContentReplaceListResponseNewString = d.Get(prefix8 + "response_new_string").(string)
		c.SlbTemplateHTTPInstanceResponseContentReplaceListResponseContentReplace = append(c.SlbTemplateHTTPInstanceResponseContentReplaceListResponseContentReplace, obj8)
	}

	SlbTemplateHTTPInstanceResponseHeaderEraseListCount := d.Get("response_header_erase_list.#").(int)
	c.SlbTemplateHTTPInstanceResponseHeaderEraseListResponseHeaderErase = make([]go_thunder.SlbTemplateHTTPInstanceResponseHeaderEraseList, 0, SlbTemplateHTTPInstanceResponseHeaderEraseListCount)

	for i := 0; i < SlbTemplateHTTPInstanceResponseHeaderEraseListCount; i++ {
		var obj9 go_thunder.SlbTemplateHTTPInstanceResponseHeaderEraseList
		prefix9 := fmt.Sprintf("response_header_erase_list.%d.", i)
		obj9.SlbTemplateHTTPInstanceResponseHeaderEraseListResponseHeaderErase = d.Get(prefix9 + "response_header_erase").(string)
		c.SlbTemplateHTTPInstanceResponseHeaderEraseListResponseHeaderErase = append(c.SlbTemplateHTTPInstanceResponseHeaderEraseListResponseHeaderErase, obj9)
	}

	SlbTemplateHTTPInstanceResponseHeaderInsertListCount := d.Get("response_header_insert_list.#").(int)
	c.SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsert = make([]go_thunder.SlbTemplateHTTPInstanceResponseHeaderInsertList, 0, SlbTemplateHTTPInstanceResponseHeaderInsertListCount)

	for i := 0; i < SlbTemplateHTTPInstanceResponseHeaderInsertListCount; i++ {
		var obj10 go_thunder.SlbTemplateHTTPInstanceResponseHeaderInsertList
		prefix10 := fmt.Sprintf("response_header_insert_list.%d.", i)
		obj10.SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsert = d.Get(prefix10 + "response_header_insert").(string)
		obj10.SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsertType = d.Get(prefix10 + "response_header_insert_type").(string)
		c.SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsert = append(c.SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsert, obj10)
	}

	c.SlbTemplateHTTPInstanceRequestTimeout = d.Get("request_timeout").(int)
	c.SlbTemplateHTTPInstanceRetryOn5Xx = d.Get("retry_on_5xx").(int)
	c.SlbTemplateHTTPInstanceRetryOn5XxVal = d.Get("retry_on_5xx_val").(int)
	c.SlbTemplateHTTPInstanceRetryOn5XxPerReq = d.Get("retry_on_5xx_per_req").(int)
	c.SlbTemplateHTTPInstanceRetryOn5XxPerReqVal = d.Get("retry_on_5xx_per_req_val").(int)
	c.SlbTemplateHTTPInstanceStrictTransactionSwitch = d.Get("strict_transaction_switch").(int)

	var obj11 go_thunder.SlbTemplateHTTPInstanceTemplate
	prefix11 := "template.0."
	obj11.SlbTemplateHTTPInstanceTemplateLogging = d.Get(prefix11 + "logging").(string)

	c.SlbTemplateHTTPInstanceTemplateLogging = obj11

	c.SlbTemplateHTTPInstanceTerm11ClientHdrConnClose = d.Get("term_11client_hdr_conn_close").(int)
	c.SlbTemplateHTTPInstancePersistOn401 = d.Get("persist_on_401").(int)
	c.SlbTemplateHTTPInstanceNum100ContWaitForReqComplete = d.Get("100_cont_wait_for_req_complete").(int)
	c.SlbTemplateHTTPInstanceUrlHashPersist = d.Get("url_hash_persist").(int)
	c.SlbTemplateHTTPInstanceUrlHashOffset = d.Get("url_hash_offset").(int)
	c.SlbTemplateHTTPInstanceUrlHashFirst = d.Get("url_hash_first").(int)
	c.SlbTemplateHTTPInstanceUrlHashLast = d.Get("url_hash_last").(int)
	c.SlbTemplateHTTPInstanceUseServerStatus = d.Get("use_server_status").(int)

	SlbTemplateHTTPInstanceUrlSwitchingCount := d.Get("url_switching.#").(int)
	c.SlbTemplateHTTPInstanceUrlSwitchingUrlSwitchingType = make([]go_thunder.SlbTemplateHTTPInstanceUrlSwitching, 0, SlbTemplateHTTPInstanceUrlSwitchingCount)

	for i := 0; i < SlbTemplateHTTPInstanceUrlSwitchingCount; i++ {
		var obj12 go_thunder.SlbTemplateHTTPInstanceUrlSwitching
		prefix12 := fmt.Sprintf("url_switching.%d.", i)
		obj12.SlbTemplateHTTPInstanceUrlSwitchingUrlSwitchingType = d.Get(prefix12 + "url_switching_type").(string)
		obj12.SlbTemplateHTTPInstanceUrlSwitchingUrlMatchString = d.Get(prefix12 + "url_match_string").(string)
		obj12.SlbTemplateHTTPInstanceUrlSwitchingUrlServiceGroup = d.Get(prefix12 + "url_service_group").(string)
		c.SlbTemplateHTTPInstanceUrlSwitchingUrlSwitchingType = append(c.SlbTemplateHTTPInstanceUrlSwitchingUrlSwitchingType, obj12)
	}

	c.SlbTemplateHTTPInstanceReqHdrWaitTime = d.Get("req_hdr_wait_time").(int)
	c.SlbTemplateHTTPInstanceReqHdrWaitTimeVal = d.Get("req_hdr_wait_time_val").(int)
	c.SlbTemplateHTTPInstanceRequestLineCaseInsensitive = d.Get("request_line_case_insensitive").(int)
	c.SlbTemplateHTTPInstanceKeepClientAlive = d.Get("keep_client_alive").(int)
	c.SlbTemplateHTTPInstanceCookieFormat = d.Get("cookie_format").(string)
	c.SlbTemplateHTTPInstancePrefix = d.Get("prefix").(string)
	c.SlbTemplateHTTPInstanceCookieSamesite = d.Get("cookie_samesite").(string)
	c.SlbTemplateHTTPInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateHTTPInstanceName = c
	return vc
}
