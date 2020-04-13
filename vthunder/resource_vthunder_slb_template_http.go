package vthunder

//vThunder resource SlbTemplateHTTP

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateHTTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateHTTPCreate,
		Update: resourceSlbTemplateHTTPUpdate,
		Read:   resourceSlbTemplateHTTPRead,
		Delete: resourceSlbTemplateHTTPDelete,
		Schema: map[string]*schema.Schema{
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
						"url_service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"url_match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"compression_auto_disable_on_high_cpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_last": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"req_hdr_wait_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_persist": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"compression_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_server_status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rd_secure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_ip_hdr_replace": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"100_cont_wait_for_req_complete": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_ip_header_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"req_hdr_wait_time_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persist_on_401": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"url_hash_first": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"strict_transaction_switch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rd_simple_loc": {
				Type:        schema.TypeString,
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
						"response_new_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"response_content_replace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"rd_resp_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"max_concurrent_streams": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bypass_sg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"term_11client_hdr_conn_close": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"insert_client_ip": {
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
			"client_port_hdr_replace": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"retry_on_5xx": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"redirect_rewrite": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redirect_secure_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"redirect_secure": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"match_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rewrite_to": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"redirect_match": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
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
			"cookie_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"compression_keep_accept_encoding_enable": {
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
			"retry_on_5xx_per_req_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_client_port_header_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"retry_on_5xx_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"url_hash_offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rd_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"host_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"host_switching_type": {
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
			"retry_on_5xx_per_req": {
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
			"failover_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"non_http_bypass": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateHTTPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateHTTP (Inside resourceSlbTemplateHTTPCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateHTTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateHTTP --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateHTTP(client.Token, data, client.Host)

		return resourceSlbTemplateHTTPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateHTTPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateHTTP (Inside resourceSlbTemplateHTTPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateHTTP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateHTTPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateHTTP   (Inside resourceSlbTemplateHTTPUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateHTTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateHTTP ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateHTTP(client.Token, name, data, client.Host)

		return resourceSlbTemplateHTTPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateHTTPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateHTTPDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateHTTP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateHTTP(d *schema.ResourceData) go_vthunder.HTTP {
	var vc go_vthunder.HTTP
	var c go_vthunder.HTTPInstance

	c.KeepClientAlive = d.Get("keep_client_alive").(int)
	c.CompressionAutoDisableOnHighCPU = d.Get("compression_auto_disable_on_high_cpu").(int)
	c.ReqHdrWaitTime = d.Get("req_hdr_wait_time").(int)

	CompressionExcludeURICount := d.Get("compression_exclude_uri.#").(int)
	c.ExcludeURI = make([]go_vthunder.CompressionExcludeURI, 0, CompressionExcludeURICount)

	for i := 0; i < CompressionExcludeURICount; i++ {
		var obj1 go_vthunder.CompressionExcludeURI
		prefix1 := fmt.Sprintf("compression_exclude_uri.%d.", i)
		obj1.ExcludeURI = d.Get(prefix1 + "exclude_uri").(string)
		c.ExcludeURI = append(c.ExcludeURI, obj1)
	}

	c.CompressionEnable = d.Get("compression_enable").(int)
	c.CompressionKeepAcceptEncoding = d.Get("compression_keep_accept_encoding").(int)
	c.FailoverURL = d.Get("failover_url").(string)

	var obj2 go_vthunder.RedirectRewrite
	prefix1 := "redirect_rewrite.0."
	obj2.RedirectSecurePort = d.Get(prefix1 + "redirect_secure_port").(int)
	obj2.RedirectSecure = d.Get(prefix1 + "redirect_secure").(int)

	MatchListCount := d.Get(prefix1 + "match_list.#").(int)
	obj2.RewriteTo = make([]go_vthunder.MatchList, 0, MatchListCount)

	for i := 0; i < MatchListCount; i++ {
		var obj2_1 go_vthunder.MatchList
		prefix2 := fmt.Sprintf(prefix1+"match_list.%d.", i)
		obj2_1.RewriteTo = d.Get(prefix2 + "rewrite_to").(string)
		obj2_1.RedirectMatch = d.Get(prefix2 + "redirect_match").(string)
		obj2.RewriteTo = append(obj2.RewriteTo, obj2_1)
	}

	c.RedirectSecurePort = obj2

	RequestHeaderEraseListCount := d.Get("request_header_erase_list.#").(int)
	c.RequestHeaderErase = make([]go_vthunder.RequestHeaderEraseList, 0, RequestHeaderEraseListCount)

	for i := 0; i < RequestHeaderEraseListCount; i++ {
		var obj3 go_vthunder.RequestHeaderEraseList
		prefix1 := fmt.Sprintf("request_header_erase_list.%d.", i)
		obj3.RequestHeaderErase = d.Get(prefix1 + "request_header_erase").(string)
		c.RequestHeaderErase = append(c.RequestHeaderErase, obj3)
	}

	c.RdPort = d.Get("rd_port").(int)

	HostSwitchingCount := d.Get("host_switching.#").(int)
	c.HostSwitchingType = make([]go_vthunder.HostSwitching, 0, HostSwitchingCount)

	for i := 0; i < HostSwitchingCount; i++ {
		var obj4 go_vthunder.HostSwitching
		prefix1 := fmt.Sprintf("host_switching.%d.", i)
		obj4.HostSwitchingType = d.Get(prefix1 + "host_switching_type").(string)
		obj4.HostServiceGroup = d.Get(prefix1 + "host_service_group").(string)
		obj4.HostMatchString = d.Get(prefix1 + "host_match_string").(string)
		c.HostSwitchingType = append(c.HostSwitchingType, obj4)
	}

	c.URLHashLast = d.Get("url_hash_last").(int)
	c.ClientIPHdrReplace = d.Get("client_ip_hdr_replace").(int)
	c.UseServerStatus = d.Get("use_server_status").(int)
	c.ReqHdrWaitTimeVal = d.Get("req_hdr_wait_time_val").(int)

	ResponseHeaderInsertListCount := d.Get("response_header_insert_list.#").(int)
	c.ResponseHeaderInsertType = make([]go_vthunder.ResponseHeaderInsertList, 0, ResponseHeaderInsertListCount)

	for i := 0; i < ResponseHeaderInsertListCount; i++ {
		var obj5 go_vthunder.ResponseHeaderInsertList
		prefix1 := fmt.Sprintf("response_header_insert_list.%d.", i)
		obj5.ResponseHeaderInsertType = d.Get(prefix1 + "response_header_insert_type").(string)
		obj5.ResponseHeaderInsert = d.Get(prefix1 + "response_header_insert").(string)
		c.ResponseHeaderInsertType = append(c.ResponseHeaderInsertType, obj5)
	}

	c.PersistOn401 = d.Get("persist_on_401").(int)
	c.Redirect = d.Get("redirect").(int)
	c.InsertClientPort = d.Get("insert_client_port").(int)
	c.RetryOn5XxPerReqVal = d.Get("retry_on_5xx_per_req_val").(int)
	c.URLHashOffset = d.Get("url_hash_offset").(int)
	c.RdSimpleLoc = d.Get("rd_simple_loc").(string)
	c.LogRetry = d.Get("log_retry").(int)
	c.NonHTTPBypass = d.Get("non_http_bypass").(int)
	c.RetryOn5XxPerReq = d.Get("retry_on_5xx_per_req").(int)
	c.InsertClientIP = d.Get("insert_client_ip").(int)

	var obj6 go_vthunder.Template2
	prefix1 = "template.0."
	obj6.Logging = d.Get(prefix1 + "logging").(string)
	c.Logging = obj6

	c.RequestTimeout = d.Get("request_timeout").(int)

	URLSwitchingCount := d.Get("url_switching.#").(int)
	c.URLServiceGroup = make([]go_vthunder.URLSwitching, 0, URLSwitchingCount)

	for i := 0; i < URLSwitchingCount; i++ {
		var obj7 go_vthunder.URLSwitching
		prefix1 := fmt.Sprintf("url_switching.%d.", i)
		obj7.URLServiceGroup = d.Get(prefix1 + "url_service_group").(string)
		obj7.URLMatchString = d.Get(prefix1 + "url_match_string").(string)
		obj7.URLSwitchingType = d.Get(prefix1 + "url_switching_type").(string)
		c.URLServiceGroup = append(c.URLServiceGroup, obj7)
	}

	c.InsertClientPortHeaderName = d.Get("insert_client_port_header_name").(string)
	c.StrictTransactionSwitch = d.Get("strict_transaction_switch").(int)

	ResponseContentReplaceListCount := d.Get("response_content_replace_list.#").(int)
	c.ResponseNewString = make([]go_vthunder.ResponseContentReplaceList, 0, ResponseContentReplaceListCount)

	for i := 0; i < ResponseContentReplaceListCount; i++ {
		var obj8 go_vthunder.ResponseContentReplaceList
		prefix1 := fmt.Sprintf("response_content_replace_list.%d.", i)
		obj8.ResponseNewString = d.Get(prefix1 + "response_new_string").(string)
		obj8.ResponseContentReplace = d.Get(prefix1 + "response_content_replace").(string)
		c.ResponseNewString = append(c.ResponseNewString, obj8)
	}

	c.One00ContWaitForReqComplete = d.Get("100_cont_wait_for_req_complete").(int)
	c.MaxConcurrentStreams = d.Get("max_concurrent_streams").(int)

	RequestHeaderInsertListCount := d.Get("request_header_insert_list.#").(int)
	c.RequestHeaderInsert = make([]go_vthunder.RequestHeaderInsertList, 0, RequestHeaderInsertListCount)

	for i := 0; i < RequestHeaderInsertListCount; i++ {
		var obj9 go_vthunder.RequestHeaderInsertList
		prefix1 := fmt.Sprintf("request_header_insert_list.%d.", i)
		obj9.RequestHeaderInsert = d.Get(prefix1 + "request_header_insert").(string)
		obj9.RequestHeaderInsertType = d.Get(prefix1 + "request_header_insert_type").(string)
		c.RequestHeaderInsert = append(c.RequestHeaderInsert, obj9)
	}

	c.CompressionMinimumContentLength = d.Get("compression_minimum_content_length").(int)
	c.CompressionLevel = d.Get("compression_level").(int)
	c.RequestLineCaseInsensitive = d.Get("request_line_case_insensitive").(int)
	c.URLHashPersist = d.Get("url_hash_persist").(int)

	ResponseHeaderEraseListCount := d.Get("response_header_erase_list.#").(int)
	c.ResponseHeaderErase = make([]go_vthunder.ResponseHeaderEraseList, 0, ResponseHeaderEraseListCount)

	for i := 0; i < ResponseHeaderEraseListCount; i++ {
		var obj10 go_vthunder.ResponseHeaderEraseList
		prefix1 := fmt.Sprintf("response_header_erase_list.%d.", i)
		obj10.ResponseHeaderErase = d.Get(prefix1 + "response_header_erase").(string)
		c.ResponseHeaderErase = append(c.ResponseHeaderErase, obj10)
	}

	c.BypassSg = d.Get("bypass_sg").(string)
	c.Name = d.Get("name").(string)
	c.RetryOn5XxVal = d.Get("retry_on_5xx_val").(int)
	c.URLHashFirst = d.Get("url_hash_first").(int)
	c.CompressionKeepAcceptEncodingEnable = d.Get("compression_keep_accept_encoding_enable").(int)
	c.UserTag = d.Get("user_tag").(string)

	CompressionContentTypeCount := d.Get("compression_content_type.#").(int)
	c.ContentType = make([]go_vthunder.CompressionContentType, 0, CompressionContentTypeCount)

	for i := 0; i < CompressionContentTypeCount; i++ {
		var obj11 go_vthunder.CompressionContentType
		prefix1 := fmt.Sprintf("compression_content_type.%d.", i)
		obj11.ContentType = d.Get(prefix1 + "content_type").(string)
		c.ContentType = append(c.ContentType, obj11)
	}

	c.ClientPortHdrReplace = d.Get("client_port_hdr_replace").(int)
	c.InsertClientIPHeaderName = d.Get("insert_client_ip_header_name").(string)
	c.RdSecure = d.Get("rd_secure").(int)
	c.RetryOn5Xx = d.Get("retry_on_5xx").(int)
	c.CookieFormat = d.Get("cookie_format").(string)
	c.Term11ClientHdrConnClose = d.Get("term_11client_hdr_conn_close").(int)

	CompressionExcludeContentTypeCount := d.Get("compression_exclude_content_type.#").(int)
	c.ExcludeContentType = make([]go_vthunder.CompressionExcludeContentType, 0, CompressionExcludeContentTypeCount)

	for i := 0; i < CompressionExcludeContentTypeCount; i++ {
		var obj12 go_vthunder.CompressionExcludeContentType
		prefix1 := fmt.Sprintf("compression_exclude_content_type.%d.", i)
		obj12.ExcludeContentType = d.Get(prefix1 + "exclude_content_type").(string)
		c.ExcludeContentType = append(c.ExcludeContentType, obj12)
	}

	c.RdRespCode = d.Get("rd_resp_code").(string)
	vc.UUID = c
	return vc
}
