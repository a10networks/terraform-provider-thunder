package thunder

//Thunder resource SlbTemplateDns

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateDns() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateDnsCreate,
		UpdateContext: resourceSlbTemplateDnsUpdate,
		ReadContext:   resourceSlbTemplateDnsRead,
		DeleteContext: resourceSlbTemplateDnsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"default_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cache_record_serving_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"remove_aa_flag": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_dns_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"drop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"max_query_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_cache_entry_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_cache_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enable_cache_sharing": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"remove_padding_to_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"add_padding_to_client": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"remove_edns_csubnet_to_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_ipv4": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"insert_ipv6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"redirect_to_tcp_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"query_id_switch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dnssec_service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"disable_rpz_attach_soa": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_logging": {
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
			"dns64": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cache": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"change_query": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"parallel_query": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"retry": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"single_response_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"timeout": {
							Type:        schema.TypeInt,
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
			"udp_retransmit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retry_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_trials": {
							Type:        schema.TypeInt,
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
			"query_type_filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_type_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_type": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"num_query_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"order": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"query_class_filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_class_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"query_class": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_class": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"num_query_class": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"rpz_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"name": {
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
						"logging": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rpz_action": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str_rpz_action": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
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
				},
			},
			"class_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"lid_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lidnum": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lid_response_rate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lid_slip_rate": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lid_window": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lid_enable_log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lid_action": {
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
							},
						},
					},
				},
			},
			"response_rate_limiting": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"filter_response_rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"slip_rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"window": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enable_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"rrl_class_list_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
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
									"lid_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lidnum": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"lid_response_rate": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"lid_slip_rate": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"lid_window": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"lid_enable_log": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"lid_action": {
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
										},
									},
								},
							},
						},
					},
				},
			},
			"local_dns_resolution": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostnames": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"local_resolver_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_resolver": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"recursive_dns_resolution": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostnames": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"csubnet_retry": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipv4_nat_pool": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_nat_pool": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"retries_per_level": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"minimal_response": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_trials": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"default_recursive": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"lookup_order": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query_type": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str_query_type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"num_query_type": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"order": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
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
				},
			},
		},
	}
}

func resourceSlbTemplateDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateDns (Inside resourceSlbTemplateDnsCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateDns(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDns --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateDns(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateDnsRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateDns (Inside resourceSlbTemplateDnsRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateDns(client.Token, name1, client.Host)
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

func resourceSlbTemplateDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateDns   (Inside resourceSlbTemplateDnsUpdate) ")
		data := dataToSlbTemplateDns(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDns ")
		err := go_thunder.PutSlbTemplateDns(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateDnsRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateDnsDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateDns(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateDns(d *schema.ResourceData) go_thunder.SlbTemplateDns {
	var vc go_thunder.SlbTemplateDns
	var c go_thunder.SlbTemplateDNSInstance
	c.SlbTemplateDNSInstanceName = d.Get("name").(string)
	c.SlbTemplateDNSInstanceDefaultPolicy = d.Get("default_policy").(string)
	c.SlbTemplateDNSInstanceCacheRecordServingPolicy = d.Get("cache_record_serving_policy").(string)
	c.SlbTemplateDNSInstanceRemoveAaFlag = d.Get("remove_aa_flag").(int)
	c.SlbTemplateDNSInstanceDisableDNSTemplate = d.Get("disable_dns_template").(int)
	c.SlbTemplateDNSInstancePeriod = d.Get("period").(int)
	c.SlbTemplateDNSInstanceDrop = d.Get("drop").(int)
	c.SlbTemplateDNSInstanceForward = d.Get("forward").(string)
	c.SlbTemplateDNSInstanceMaxQueryLength = d.Get("max_query_length").(int)
	c.SlbTemplateDNSInstanceMaxCacheEntrySize = d.Get("max_cache_entry_size").(int)
	c.SlbTemplateDNSInstanceMaxCacheSize = d.Get("max_cache_size").(int)
	c.SlbTemplateDNSInstanceEnableCacheSharing = d.Get("enable_cache_sharing").(int)
	c.SlbTemplateDNSInstanceRemovePaddingToServer = d.Get("remove_padding_to_server").(int)
	c.SlbTemplateDNSInstanceAddPaddingToClient = d.Get("add_padding_to_client").(string)
	c.SlbTemplateDNSInstanceRemoveEdnsCsubnetToServer = d.Get("remove_edns_csubnet_to_server").(int)
	c.SlbTemplateDNSInstanceInsertIpv4 = d.Get("insert_ipv4").(int)
	c.SlbTemplateDNSInstanceInsertIpv6 = d.Get("insert_ipv6").(int)
	c.SlbTemplateDNSInstanceRedirectToTCPPort = d.Get("redirect_to_tcp_port").(int)
	c.SlbTemplateDNSInstanceQueryIDSwitch = d.Get("query_id_switch").(int)
	c.SlbTemplateDNSInstanceDnssecServiceGroup = d.Get("dnssec_service_group").(string)
	c.SlbTemplateDNSInstanceDisableRpzAttachSoa = d.Get("disable_rpz_attach_soa").(int)
	c.SlbTemplateDNSInstanceDNSLogging = d.Get("dns_logging").(string)
	c.SlbTemplateDNSInstanceUserTag = d.Get("user_tag").(string)

	var obj1 go_thunder.SlbTemplateDNSInstanceDNS64
	prefix1 := "dns64.0."
	obj1.SlbTemplateDNSInstanceDNS64Enable = d.Get(prefix1 + "enable").(int)
	obj1.SlbTemplateDNSInstanceDNS64Cache = d.Get(prefix1 + "cache").(int)
	obj1.SlbTemplateDNSInstanceDNS64ChangeQuery = d.Get(prefix1 + "change_query").(int)
	obj1.SlbTemplateDNSInstanceDNS64ParallelQuery = d.Get(prefix1 + "parallel_query").(int)
	obj1.SlbTemplateDNSInstanceDNS64Retry = d.Get(prefix1 + "retry").(int)
	obj1.SlbTemplateDNSInstanceDNS64SingleResponseDisable = d.Get(prefix1 + "single_response_disable").(int)
	obj1.SlbTemplateDNSInstanceDNS64Timeout = d.Get(prefix1 + "timeout").(int)

	c.SlbTemplateDNSInstanceDNS64Enable = obj1

	var obj2 go_thunder.SlbTemplateDNSInstanceUDPRetransmit
	prefix2 := "udp_retransmit.0."
	obj2.SlbTemplateDNSInstanceUDPRetransmitRetryInterval = d.Get(prefix2 + "retry_interval").(int)
	obj2.SlbTemplateDNSInstanceUDPRetransmitMaxTrials = d.Get(prefix2 + "max_trials").(int)

	c.SlbTemplateDNSInstanceUDPRetransmitRetryInterval = obj2

	var obj3 go_thunder.SlbTemplateDNSInstanceQueryTypeFilter
	prefix3 := "query_type_filter.0."
	obj3.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeAction = d.Get(prefix3 + "query_type_action").(string)

	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeCount := d.Get(prefix3 + "query_type.#").(int)
	obj3.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeStrQueryType = make([]go_thunder.SlbTemplateDNSInstanceQueryTypeFilterQueryType, 0, SlbTemplateDNSInstanceQueryTypeFilterQueryTypeCount)

	for i := 0; i < SlbTemplateDNSInstanceQueryTypeFilterQueryTypeCount; i++ {
		var obj3_1 go_thunder.SlbTemplateDNSInstanceQueryTypeFilterQueryType
		prefix3_1 := prefix3 + fmt.Sprintf("query_type.%d.", i)
		obj3_1.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeStrQueryType = d.Get(prefix3_1 + "str_query_type").(string)
		obj3_1.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeNumQueryType = d.Get(prefix3_1 + "num_query_type").(int)
		obj3_1.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeOrder = d.Get(prefix3_1 + "order").(string)
		obj3.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeStrQueryType = append(obj3.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeStrQueryType, obj3_1)
	}

	c.SlbTemplateDNSInstanceQueryTypeFilterQueryTypeAction = obj3

	var obj4 go_thunder.SlbTemplateDNSInstanceQueryClassFilter
	prefix4 := "query_class_filter.0."
	obj4.SlbTemplateDNSInstanceQueryClassFilterQueryClassAction = d.Get(prefix4 + "query_class_action").(string)

	SlbTemplateDNSInstanceQueryClassFilterQueryClassCount := d.Get(prefix4 + "query_class.#").(int)
	obj4.SlbTemplateDNSInstanceQueryClassFilterQueryClassStrQueryClass = make([]go_thunder.SlbTemplateDNSInstanceQueryClassFilterQueryClass, 0, SlbTemplateDNSInstanceQueryClassFilterQueryClassCount)

	for i := 0; i < SlbTemplateDNSInstanceQueryClassFilterQueryClassCount; i++ {
		var obj4_1 go_thunder.SlbTemplateDNSInstanceQueryClassFilterQueryClass
		prefix4_1 := prefix4 + fmt.Sprintf("query_class.%d.", i)
		obj4_1.SlbTemplateDNSInstanceQueryClassFilterQueryClassStrQueryClass = d.Get(prefix4_1 + "str_query_class").(string)
		obj4_1.SlbTemplateDNSInstanceQueryClassFilterQueryClassNumQueryClass = d.Get(prefix4_1 + "num_query_class").(int)
		obj4.SlbTemplateDNSInstanceQueryClassFilterQueryClassStrQueryClass = append(obj4.SlbTemplateDNSInstanceQueryClassFilterQueryClassStrQueryClass, obj4_1)
	}

	c.SlbTemplateDNSInstanceQueryClassFilterQueryClassAction = obj4

	SlbTemplateDNSInstanceRpzListCount := d.Get("rpz_list.#").(int)
	c.SlbTemplateDNSInstanceRpzListSeqID = make([]go_thunder.SlbTemplateDNSInstanceRpzList, 0, SlbTemplateDNSInstanceRpzListCount)

	for i := 0; i < SlbTemplateDNSInstanceRpzListCount; i++ {
		var obj5 go_thunder.SlbTemplateDNSInstanceRpzList
		prefix5 := fmt.Sprintf("rpz_list.%d.", i)
		obj5.SlbTemplateDNSInstanceRpzListSeqID = d.Get(prefix5 + "seq_id").(int)
		obj5.SlbTemplateDNSInstanceRpzListName = d.Get(prefix5 + "name").(string)
		obj5.SlbTemplateDNSInstanceRpzListUserTag = d.Get(prefix5 + "user_tag").(string)

		var obj5_1 go_thunder.SlbTemplateDNSInstanceRpzListLogging
		prefix5_1 := prefix5 + "logging.0."
		obj5_1.SlbTemplateDNSInstanceRpzListLoggingEnable = d.Get(prefix5_1 + "enable").(int)

		SlbTemplateDNSInstanceRpzListLoggingRpzActionCount := d.Get(prefix5_1 + "rpz_action.#").(int)
		obj5_1.SlbTemplateDNSInstanceRpzListLoggingRpzActionStrRpzAction = make([]go_thunder.SlbTemplateDNSInstanceRpzListLoggingRpzAction, 0, SlbTemplateDNSInstanceRpzListLoggingRpzActionCount)

		for i := 0; i < SlbTemplateDNSInstanceRpzListLoggingRpzActionCount; i++ {
			var obj5_1_1 go_thunder.SlbTemplateDNSInstanceRpzListLoggingRpzAction
			prefix5_1_1 := prefix5_1 + fmt.Sprintf("rpz_action.%d.", i)
			obj5_1_1.SlbTemplateDNSInstanceRpzListLoggingRpzActionStrRpzAction = d.Get(prefix5_1_1 + "str_rpz_action").(string)
			obj5_1.SlbTemplateDNSInstanceRpzListLoggingRpzActionStrRpzAction = append(obj5_1.SlbTemplateDNSInstanceRpzListLoggingRpzActionStrRpzAction, obj5_1_1)
		}

		obj5.SlbTemplateDNSInstanceRpzListLoggingEnable = obj5_1

		c.SlbTemplateDNSInstanceRpzListSeqID = append(c.SlbTemplateDNSInstanceRpzListSeqID, obj5)
	}

	var obj6 go_thunder.SlbTemplateDNSInstanceClassList
	prefix6 := "class_list.0."
	obj6.SlbTemplateDNSInstanceClassListName = d.Get(prefix6 + "name").(string)

	SlbTemplateDNSInstanceClassListLidListCount := d.Get(prefix6 + "lid_list.#").(int)
	obj6.SlbTemplateDNSInstanceClassListLidListLidnum = make([]go_thunder.SlbTemplateDNSInstanceClassListLidList, 0, SlbTemplateDNSInstanceClassListLidListCount)

	for i := 0; i < SlbTemplateDNSInstanceClassListLidListCount; i++ {
		var obj6_1 go_thunder.SlbTemplateDNSInstanceClassListLidList
		prefix6_1 := prefix6 + fmt.Sprintf("lid_list.%d.", i)
		obj6_1.SlbTemplateDNSInstanceClassListLidListLidnum = d.Get(prefix6_1 + "lidnum").(int)
		obj6_1.SlbTemplateDNSInstanceClassListLidListLidResponseRate = d.Get(prefix6_1 + "lid_response_rate").(int)
		obj6_1.SlbTemplateDNSInstanceClassListLidListLidSlipRate = d.Get(prefix6_1 + "lid_slip_rate").(int)
		obj6_1.SlbTemplateDNSInstanceClassListLidListLidWindow = d.Get(prefix6_1 + "lid_window").(int)
		obj6_1.SlbTemplateDNSInstanceClassListLidListLidEnableLog = d.Get(prefix6_1 + "lid_enable_log").(int)
		obj6_1.SlbTemplateDNSInstanceClassListLidListLidAction = d.Get(prefix6_1 + "lid_action").(string)
		obj6_1.SlbTemplateDNSInstanceClassListLidListUserTag = d.Get(prefix6_1 + "user_tag").(string)
		obj6.SlbTemplateDNSInstanceClassListLidListLidnum = append(obj6.SlbTemplateDNSInstanceClassListLidListLidnum, obj6_1)
	}

	c.SlbTemplateDNSInstanceClassListName = obj6

	var obj7 go_thunder.SlbTemplateDNSInstanceResponseRateLimiting
	prefix7 := "response_rate_limiting.0."
	obj7.SlbTemplateDNSInstanceResponseRateLimitingResponseRate = d.Get(prefix7 + "response_rate").(int)
	obj7.SlbTemplateDNSInstanceResponseRateLimitingFilterResponseRate = d.Get(prefix7 + "filter_response_rate").(int)
	obj7.SlbTemplateDNSInstanceResponseRateLimitingSlipRate = d.Get(prefix7 + "slip_rate").(int)
	obj7.SlbTemplateDNSInstanceResponseRateLimitingWindow = d.Get(prefix7 + "window").(int)
	obj7.SlbTemplateDNSInstanceResponseRateLimitingEnableLog = d.Get(prefix7 + "enable_log").(int)
	obj7.SlbTemplateDNSInstanceResponseRateLimitingAction = d.Get(prefix7 + "action").(string)

	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListCount := d.Get(prefix7 + "rrl_class_list_list.#").(int)
	obj7.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListName = make([]go_thunder.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListList, 0, SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListCount)

	for i := 0; i < SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListCount; i++ {
		var obj7_1 go_thunder.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListList
		prefix7_1 := prefix7 + fmt.Sprintf("rrl_class_list_list.%d.", i)
		obj7_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListName = d.Get(prefix7_1 + "name").(string)
		obj7_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListUserTag = d.Get(prefix7_1 + "user_tag").(string)

		SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListCount := d.Get(prefix7_1 + "lid_list.#").(int)
		obj7_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidnum = make([]go_thunder.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidList, 0, SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListCount)

		for i := 0; i < SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListCount; i++ {
			var obj7_1_1 go_thunder.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidList
			prefix7_1_1 := prefix7_1 + fmt.Sprintf("lid_list.%d.", i)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidnum = d.Get(prefix7_1_1 + "lidnum").(int)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidResponseRate = d.Get(prefix7_1_1 + "lid_response_rate").(int)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidSlipRate = d.Get(prefix7_1_1 + "lid_slip_rate").(int)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidWindow = d.Get(prefix7_1_1 + "lid_window").(int)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidEnableLog = d.Get(prefix7_1_1 + "lid_enable_log").(int)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidAction = d.Get(prefix7_1_1 + "lid_action").(string)
			obj7_1_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListUserTag = d.Get(prefix7_1_1 + "user_tag").(string)
			obj7_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidnum = append(obj7_1.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidnum, obj7_1_1)
		}

		obj7.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListName = append(obj7.SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListName, obj7_1)
	}

	c.SlbTemplateDNSInstanceResponseRateLimitingResponseRate = obj7

	var obj8 go_thunder.SlbTemplateDNSInstanceLocalDNSResolution
	prefix8 := "local_dns_resolution.0."

	SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgCount := d.Get(prefix8 + "host_list_cfg.#").(int)
	obj8.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgHostnames = make([]go_thunder.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfg, 0, SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgCount)

	for i := 0; i < SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgCount; i++ {
		var obj8_1 go_thunder.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfg
		prefix8_1 := prefix8 + fmt.Sprintf("host_list_cfg.%d.", i)
		obj8_1.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgHostnames = d.Get(prefix8_1 + "hostnames").(string)
		obj8.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgHostnames = append(obj8.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgHostnames, obj8_1)
	}

	SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgCount := d.Get(prefix8 + "local_resolver_cfg.#").(int)
	obj8.SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgLocalResolver = make([]go_thunder.SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfg, 0, SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgCount)

	for i := 0; i < SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgCount; i++ {
		var obj8_2 go_thunder.SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfg
		prefix8_2 := prefix8 + fmt.Sprintf("local_resolver_cfg.%d.", i)
		obj8_2.SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgLocalResolver = d.Get(prefix8_2 + "local_resolver").(string)
		obj8.SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgLocalResolver = append(obj8.SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgLocalResolver, obj8_2)
	}

	c.SlbTemplateDNSInstanceLocalDNSResolutionHostListCfg = obj8

	var obj9 go_thunder.SlbTemplateDNSInstanceRecursiveDNSResolution
	prefix9 := "recursive_dns_resolution.0."

	SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgCount := d.Get(prefix9 + "host_list_cfg.#").(int)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgHostnames = make([]go_thunder.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfg, 0, SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgCount)

	for i := 0; i < SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgCount; i++ {
		var obj9_1 go_thunder.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfg
		prefix9_1 := prefix9 + fmt.Sprintf("host_list_cfg.%d.", i)
		obj9_1.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgHostnames = d.Get(prefix9_1 + "hostnames").(string)
		obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgHostnames = append(obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgHostnames, obj9_1)
	}

	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionCsubnetRetry = d.Get(prefix9 + "csubnet_retry").(int)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionIpv4NatPool = d.Get(prefix9 + "ipv4_nat_pool").(string)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionIpv6NatPool = d.Get(prefix9 + "ipv6_nat_pool").(string)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionRetriesPerLevel = d.Get(prefix9 + "retries_per_level").(int)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionMinimalResponse = d.Get(prefix9 + "minimal_response").(int)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionMaxTrials = d.Get(prefix9 + "max_trials").(int)
	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionDefaultRecursive = d.Get(prefix9 + "default_recursive").(int)

	var obj9_2 go_thunder.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrder
	prefix9_2 := prefix9 + "lookup_order.0."

	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeCount := d.Get(prefix9_2 + "query_type.#").(int)
	obj9_2.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeStrQueryType = make([]go_thunder.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryType, 0, SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeCount)

	for i := 0; i < SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeCount; i++ {
		var obj9_2_1 go_thunder.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryType
		prefix9_2_1 := prefix9_2 + fmt.Sprintf("query_type.%d.", i)
		obj9_2_1.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeStrQueryType = d.Get(prefix9_2_1 + "str_query_type").(string)
		obj9_2_1.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeNumQueryType = d.Get(prefix9_2_1 + "num_query_type").(int)
		obj9_2_1.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeOrder = d.Get(prefix9_2_1 + "order").(string)
		obj9_2.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeStrQueryType = append(obj9_2.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeStrQueryType, obj9_2_1)
	}

	obj9.SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryType = obj9_2

	c.SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfg = obj9

	vc.SlbTemplateDNSInstanceName = c
	return vc
}
