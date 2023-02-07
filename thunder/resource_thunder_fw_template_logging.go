package thunder

//Thunder resource FwTemplateLogging

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceFwTemplateLogging() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFwTemplateLoggingCreate,
		UpdateContext: resourceFwTemplateLoggingUpdate,
		ReadContext:   resourceFwTemplateLoggingRead,
		DeleteContext: resourceFwTemplateLoggingDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"resolution": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"include_dest_fqdn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"merged_style": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_requests": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"include_radius_attribute": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"attr_event": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"no_quote": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"framed_ipv6_prefix": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"prefix_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"insert_if_not_existing": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"zero_in_custom_attr": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"include_http": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_header": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"max_length": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"custom_header_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"custom_max_length": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"l4_session_info": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"method": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"request_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"file_extension": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"rule": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_http_requests": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dest_port": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dest_port_number": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"include_byte_count": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"log_every_http_request": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"max_url_len": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"include_all_headers": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable_sequence_check": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"facility": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"severity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_group": {
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
			"source_address": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwTemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating FwTemplateLogging (Inside resourceFwTemplateLoggingCreate) ")
		name1 := d.Get("name").(string)
		data := dataToFwTemplateLogging(d)
		logger.Println("[INFO] received formatted data from method data to FwTemplateLogging --")
		d.SetId(name1)
		err := go_thunder.PostFwTemplateLogging(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTemplateLoggingRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwTemplateLogging (Inside resourceFwTemplateLoggingRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetFwTemplateLogging(client.Token, name1, client.Host)
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

func resourceFwTemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying FwTemplateLogging   (Inside resourceFwTemplateLoggingUpdate) ")
		data := dataToFwTemplateLogging(d)
		logger.Println("[INFO] received formatted data from method data to FwTemplateLogging ")
		err := go_thunder.PutFwTemplateLogging(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceFwTemplateLoggingRead(ctx, d, meta)

	}
	return diags
}

func resourceFwTemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceFwTemplateLoggingDelete) " + name1)
		err := go_thunder.DeleteFwTemplateLogging(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToFwTemplateLogging(d *schema.ResourceData) go_thunder.FwTemplateLogging {
	var vc go_thunder.FwTemplateLogging
	var c go_thunder.FwTemplateLoggingInstance
	c.FwTemplateLoggingInstanceName = d.Get("name").(string)
	c.FwTemplateLoggingInstanceResolution = d.Get("resolution").(string)
	c.FwTemplateLoggingInstanceIncludeDestFqdn = d.Get("include_dest_fqdn").(int)
	c.FwTemplateLoggingInstanceMergedStyle = d.Get("merged_style").(int)

	var obj1 go_thunder.FwTemplateLoggingInstanceLog
	prefix1 := "log.0."
	obj1.FwTemplateLoggingInstanceLogHTTPRequests = d.Get(prefix1 + "http_requests").(string)

	c.FwTemplateLoggingInstanceLogHTTPRequests = obj1

	var obj2 go_thunder.FwTemplateLoggingInstanceIncludeRadiusAttribute
	prefix2 := "include_radius_attribute.0."

	FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgCount := d.Get(prefix2 + "attr_cfg.#").(int)
	obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttr = make([]go_thunder.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfg, 0, FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgCount)

	for i := 0; i < FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgCount; i++ {
		var obj2_1 go_thunder.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfg
		prefix2_1 := prefix2 + fmt.Sprintf("attr_cfg.%d.", i)
		obj2_1.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttr = d.Get(prefix2_1 + "attr").(string)
		obj2_1.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttrEvent = d.Get(prefix2_1 + "attr_event").(string)
		obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttr = append(obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttr, obj2_1)
	}

	obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeNoQuote = d.Get(prefix2 + "no_quote").(int)
	obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeFramedIpv6Prefix = d.Get(prefix2 + "framed_ipv6_prefix").(int)
	obj2.FwTemplateLoggingInstanceIncludeRadiusAttributePrefixLength = d.Get(prefix2 + "prefix_length").(string)
	obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeInsertIfNotExisting = d.Get(prefix2 + "insert_if_not_existing").(int)
	obj2.FwTemplateLoggingInstanceIncludeRadiusAttributeZeroInCustomAttr = d.Get(prefix2 + "zero_in_custom_attr").(int)

	c.FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfg = obj2

	var obj3 go_thunder.FwTemplateLoggingInstanceIncludeHTTP
	prefix3 := "include_http.0."

	FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCount := d.Get(prefix3 + "header_cfg.#").(int)
	obj3.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgHTTPHeader = make([]go_thunder.FwTemplateLoggingInstanceIncludeHTTPHeaderCfg, 0, FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCount)

	for i := 0; i < FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCount; i++ {
		var obj3_1 go_thunder.FwTemplateLoggingInstanceIncludeHTTPHeaderCfg
		prefix3_1 := prefix3 + fmt.Sprintf("header_cfg.%d.", i)
		obj3_1.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgHTTPHeader = d.Get(prefix3_1 + "http_header").(string)
		obj3_1.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgMaxLength = d.Get(prefix3_1 + "max_length").(int)
		obj3_1.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCustomHeaderName = d.Get(prefix3_1 + "custom_header_name").(string)
		obj3_1.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCustomMaxLength = d.Get(prefix3_1 + "custom_max_length").(int)
		obj3.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgHTTPHeader = append(obj3.FwTemplateLoggingInstanceIncludeHTTPHeaderCfgHTTPHeader, obj3_1)
	}

	obj3.FwTemplateLoggingInstanceIncludeHTTPL4SessionInfo = d.Get(prefix3 + "l4_session_info").(int)
	obj3.FwTemplateLoggingInstanceIncludeHTTPMethod = d.Get(prefix3 + "method").(int)
	obj3.FwTemplateLoggingInstanceIncludeHTTPRequestNumber = d.Get(prefix3 + "request_number").(int)
	obj3.FwTemplateLoggingInstanceIncludeHTTPFileExtension = d.Get(prefix3 + "file_extension").(int)

	c.FwTemplateLoggingInstanceIncludeHTTPHeaderCfg = obj3

	var obj4 go_thunder.FwTemplateLoggingInstanceRule
	prefix4 := "rule.0."

	var obj4_1 go_thunder.FwTemplateLoggingInstanceRuleRuleHTTPRequests
	prefix4_1 := prefix4 + "rule_http_requests.0."

	FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortCount := d.Get(prefix4_1 + "dest_port.#").(int)
	obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortDestPortNumber = make([]go_thunder.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPort, 0, FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortCount)

	for i := 0; i < FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortCount; i++ {
		var obj4_1_1 go_thunder.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPort
		prefix4_1_1 := prefix4_1 + fmt.Sprintf("dest_port.%d.", i)
		obj4_1_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortDestPortNumber = d.Get(prefix4_1_1 + "dest_port_number").(int)
		obj4_1_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortIncludeByteCount = d.Get(prefix4_1_1 + "include_byte_count").(int)
		obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortDestPortNumber = append(obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortDestPortNumber, obj4_1_1)
	}

	obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsLogEveryHTTPRequest = d.Get(prefix4_1 + "log_every_http_request").(int)
	obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsMaxURLLen = d.Get(prefix4_1 + "max_url_len").(int)
	obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsIncludeAllHeaders = d.Get(prefix4_1 + "include_all_headers").(int)
	obj4_1.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDisableSequenceCheck = d.Get(prefix4_1 + "disable_sequence_check").(int)

	obj4.FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPort = obj4_1

	c.FwTemplateLoggingInstanceRuleRuleHTTPRequests = obj4

	c.FwTemplateLoggingInstanceFacility = d.Get("facility").(string)
	c.FwTemplateLoggingInstanceSeverity = d.Get("severity").(string)
	c.FwTemplateLoggingInstanceFormat = d.Get("format").(string)
	c.FwTemplateLoggingInstanceServiceGroup = d.Get("service_group").(string)
	c.FwTemplateLoggingInstanceUserTag = d.Get("user_tag").(string)

	var obj5 go_thunder.FwTemplateLoggingInstanceSourceAddress
	prefix5 := "source_address.0."
	obj5.FwTemplateLoggingInstanceSourceAddressIP = d.Get(prefix5 + "ip").(string)
	obj5.FwTemplateLoggingInstanceSourceAddressIpv6 = d.Get(prefix5 + "ipv6").(string)

	c.FwTemplateLoggingInstanceSourceAddressIP = obj5

	vc.FwTemplateLoggingInstanceName = c
	return vc
}
