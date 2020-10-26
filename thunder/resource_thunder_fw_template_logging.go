package thunder

//Thunder resource FwTemplateLogging

import (
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFwTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwTemplateLoggingCreate,
		Update: resourceFwTemplateLoggingUpdate,
		Read:   resourceFwTemplateLoggingRead,
		Delete: resourceFwTemplateLoggingDelete,
		Schema: map[string]*schema.Schema{
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
									"custom_max_length": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
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
								},
							},
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
						"method": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l4_session_info": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"merged_style": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"format": {
				Type:        schema.TypeString,
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
			"severity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"include_dest_fqdn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"facility": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"include_radius_attribute": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"no_quote": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"attr_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attr_event": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"attr": {
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
									"log_every_http_request": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable_sequence_check": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"include_all_headers": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dest_port": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"include_byte_count": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dest_port_number": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"max_url_len": {
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
			"service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"resolution": {
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
	}
}

func resourceFwTemplateLoggingCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTemplateLogging (Inside resourceFwTemplateLoggingCreate) ")
		name := d.Get("name").(string)
		data := dataToFwTemplateLogging(d)
		logger.Println("[INFO] received formatted data from method data to FwTemplateLogging --")
		d.SetId(name)
		go_thunder.PostFwTemplateLogging(client.Token, data, client.Host)

		return resourceFwTemplateLoggingRead(d, meta)

	}
	return nil
}

func resourceFwTemplateLoggingRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwTemplateLogging (Inside resourceFwTemplateLoggingRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwTemplateLogging(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwTemplateLoggingUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTemplateLoggingRead(d, meta)
}

func resourceFwTemplateLoggingDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTemplateLoggingRead(d, meta)
}
func dataToFwTemplateLogging(d *schema.ResourceData) go_thunder.FwTemplateLogging {
	var vc go_thunder.FwTemplateLogging
	var c go_thunder.FwTemplateLoggingInstance

	var obj1 go_thunder.FwTemplateLoggingIncludeHTTP
	prefix := "include_http.0."

	HeaderCfgCount := d.Get(prefix + "header_cfg.#").(int)
	obj1.CustomMaxLength = make([]go_thunder.FwTemplateLoggingHeaderCfg, 0, HeaderCfgCount)

	for i := 0; i < HeaderCfgCount; i++ {
		var obj1_1 go_thunder.FwTemplateLoggingHeaderCfg
		prefix1 := prefix + fmt.Sprintf("header_cfg.%d.", i)
		obj1_1.CustomMaxLength = d.Get(prefix1 + "custom_max_length").(int)
		obj1_1.HTTPHeader = d.Get(prefix1 + "http_header").(string)
		obj1_1.MaxLength = d.Get(prefix1 + "max_length").(int)
		obj1_1.CustomHeaderName = d.Get(prefix1 + "custom_header_name").(string)
		obj1.CustomMaxLength = append(obj1.CustomMaxLength, obj1_1)
	}

	obj1.RequestNumber = d.Get(prefix + "request_number").(int)
	obj1.FileExtension = d.Get(prefix + "file_extension").(int)
	obj1.Method = d.Get(prefix + "method").(int)
	obj1.L4SessionInfo = d.Get(prefix + "l4_session_info").(int)

	c.FileExtension = obj1

	c.MergedStyle = d.Get("merged_style").(int)
	c.Name = d.Get("name").(string)

	var obj1_2 go_thunder.FwTemplateLoggingSourceAddress
	prefix2 := "source_address.0."
	obj1_2.IP = d.Get(prefix2 + "ip").(string)
	obj1_2.Ipv6 = d.Get(prefix2 + "ipv6").(string)

	c.IP = obj1_2

	c.Format = d.Get("format").(string)

	var obj1_3 go_thunder.FwTemplateLoggingLog
	prefix3 := "log.0."
	obj1_3.HTTPRequests = d.Get(prefix3 + "http_requests").(string)

	c.HTTPRequests = obj1_3

	c.Severity = d.Get("severity").(string)
	c.IncludeDestFqdn = d.Get("include_dest_fqdn").(int)
	c.Facility = d.Get("facility").(string)

	var obj1_4 go_thunder.FwTemplateLoggingIncludeRadiusAttribute
	prefix4 := "include_radius_attribute.0."
	obj1_4.FramedIpv6Prefix = d.Get(prefix4 + "framed_ipv6_prefix").(int)
	obj1_4.PrefixLength = d.Get(prefix4 + "prefix_length").(string)
	obj1_4.InsertIfNotExisting = d.Get(prefix4 + "insert_if_not_existing").(int)
	obj1_4.ZeroInCustomAttr = d.Get(prefix4 + "zero_in_custom_attr").(int)
	obj1_4.NoQuote = d.Get(prefix4 + "no_quote").(int)

	AttrCfgCount := d.Get(prefix4 + "attr_cfg.#").(int)
	obj1_4.AttrEvent = make([]go_thunder.FwTemplateLoggingAttrCfg, 0, AttrCfgCount)

	for i := 0; i < AttrCfgCount; i++ {
		var obj1_4_1 go_thunder.FwTemplateLoggingAttrCfg
		prefix4_1 := prefix4 + fmt.Sprintf("attr_cfg.%d.", i)
		obj1_4_1.AttrEvent = d.Get(prefix4_1 + "attr_event").(string)
		obj1_4_1.Attr = d.Get(prefix4_1 + "attr").(string)
		obj1_4.AttrEvent = append(obj1_4.AttrEvent, obj1_4_1)
	}

	c.FramedIpv6Prefix = obj1_4

	var obj1_5 go_thunder.FwTemplateLoggingRule
	prefix5 := "rule.0."

	var obj1_5_1 go_thunder.FwTemplateLoggingRuleHTTPRequests
	prefix5_1 := prefix5 + "rule_http_requests.0."
	obj1_5_1.LogEveryHTTPRequest = d.Get(prefix5_1 + "log_every_http_request").(int)
	obj1_5_1.DisableSequenceCheck = d.Get(prefix5_1 + "disable_sequence_check").(int)
	obj1_5_1.IncludeAllHeaders = d.Get(prefix5_1 + "include_all_headers").(int)

	DestPortCount := d.Get(prefix5_1 + "dest_port.#").(int)
	obj1_5_1.IncludeByteCount = make([]go_thunder.FwTemplateLoggingDestPort, 0, DestPortCount)

	for i := 0; i < DestPortCount; i++ {
		var obj1_5_1_1 go_thunder.FwTemplateLoggingDestPort
		prefix5_1_1 := prefix5_1 + fmt.Sprintf("dest_port.%d.", i)
		obj1_5_1_1.IncludeByteCount = d.Get(prefix5_1_1 + "include_byte_count").(int)
		obj1_5_1_1.DestPortNumber = d.Get(prefix5_1_1 + "dest_port_number").(int)
		obj1_5_1.IncludeByteCount = append(obj1_5_1.IncludeByteCount, obj1_5_1_1)
	}

	obj1_5_1.MaxURLLen = d.Get(prefix5_1 + "max_url_len").(int)

	obj1_5.IncludeByteCount = obj1_5_1

	c.IncludeByteCount = obj1_5

	c.ServiceGroup = d.Get("service_group").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Resolution = d.Get("resolution").(string)

	vc.UUID = c
	return vc
}
