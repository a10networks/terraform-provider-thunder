package thunder

//Thunder resource SlbTemplatePolicy

import (
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplatePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplatePolicyCreate,
		Update: resourceSlbTemplatePolicyUpdate,
		Read:   resourceSlbTemplatePolicyRead,
		Delete: resourceSlbTemplatePolicyDelete,
		Schema: map[string]*schema.Schema{
			"over_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"full_domain_tree": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_list_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pbslb_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pbslb_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"logging_drp_rst": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bw_list_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
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
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_destination_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"over_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"overlap": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"over_limit_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"share": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"forward_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"no_client_conn_reuse": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"source_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"match_authorize_policy": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"match_any": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"destination": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"dest_class_list": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"sampling_enable": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "",
																		},
																	},
																},
															},
															"type": {
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
												"any": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"sampling_enable": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
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
												"web_category_list_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"web_category_list": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"action": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"sampling_enable": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"counters1": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "",
																		},
																	},
																},
															},
															"type": {
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
										},
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"sampling_enable": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"match_class_list": {
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
						"local_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drop_redirect_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"real_sg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"sampling_enable": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"counters1": {
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
									"drop_message": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"fall_back_snat": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"http_status_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"fall_back": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"forward_snat": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"fake_sg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"action1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"drop_response_code": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"acos_event_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"require_web_category": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"filtering": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssli_url_filtering": {
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
						"san_filtering": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssli_url_filtering_san": {
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
			"over_limit_lockup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
						"header_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"lid_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_pbslb_logging": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"conn_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
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
												"prefix": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"exclusive_answer": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"direct_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"bw_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"response_code_rate_limit": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code_range_start": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"period": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"code_range_end": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"threshold": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"lidnum": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"request_per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"request_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"over_limit_action": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_fail": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_action_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"request_rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_service_group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"lockout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"conn_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_action_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"bw_per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_pbslb_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"action_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"conn_per": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"direct_logging_drp_rst": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"client_ip_l3_dest": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"client_ip_l7_header": {
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
			"bw_list_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplatePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplatePolicy (Inside resourceSlbTemplatePolicyCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplatePolicy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplatePolicy --")
		d.SetId(name)
		go_thunder.PostSlbTemplatePolicy(client.Token, data, client.Host)

		return resourceSlbTemplatePolicyRead(d, meta)

	}
	return nil
}

func resourceSlbTemplatePolicyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplatePolicy (Inside resourceSlbTemplatePolicyRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplatePolicy(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplatePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplatePolicy   (Inside resourceSlbTemplatePolicyUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplatePolicy(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplatePolicy ")
		d.SetId(name)
		go_thunder.PutSlbTemplatePolicy(client.Token, name, data, client.Host)

		return resourceSlbTemplatePolicyRead(d, meta)

	}
	return nil
}

func resourceSlbTemplatePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplatePolicyDelete) " + name)
		err := go_thunder.DeleteSlbTemplatePolicy(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplatePolicy(d *schema.ResourceData) go_thunder.Policy {
	var vc go_thunder.Policy
	var c go_thunder.PolicyInstance

	var obj1 go_thunder.ForwardPolicy
	prefix1 := "forward_policy.0."

	FilteringCount := d.Get(prefix1 + "filtering.#").(int)
	obj1.SsliURLFiltering = make([]go_thunder.Filtering, 0, FilteringCount)

	for i := 0; i < FilteringCount; i++ {
		var obj1_1 go_thunder.Filtering
		prefix2 := fmt.Sprintf(prefix1+"filtering.%d.", i)
		obj1_1.SsliURLFiltering = d.Get(prefix2 + "ssli_url_filtering").(string)
		obj1.SsliURLFiltering = append(obj1.SsliURLFiltering, obj1_1)
	}

	obj1.LocalLogging = d.Get(prefix1 + "local_logging").(int)

	SanFilteringCount := d.Get(prefix1 + "san_filtering.#").(int)
	obj1.SsliURLFilteringSan = make([]go_thunder.SanFiltering, 0, SanFilteringCount)

	for i := 0; i < SanFilteringCount; i++ {
		var obj1_2 go_thunder.SanFiltering
		prefix2 := fmt.Sprintf(prefix1+"san_filtering.%d.", i)
		obj1_2.SsliURLFilteringSan = d.Get(prefix2 + "ssli_url_filtering_san").(string)
		obj1.SsliURLFilteringSan = append(obj1.SsliURLFilteringSan, obj1_2)
	}

	ActionListCount := d.Get(prefix1 + "action_list.#").(int)
	obj1.Log = make([]go_thunder.ActionList, 0, ActionListCount)

	for i := 0; i < ActionListCount; i++ {
		var obj1_3 go_thunder.ActionList
		prefix2 := fmt.Sprintf("forward_policy.0.action_list.%d.", i)
		obj1_3.Log = d.Get(prefix2 + "log").(int)
		obj1_3.HTTPStatusCode = d.Get(prefix2 + "http_status_code").(string)
		obj1_3.ForwardSnat = d.Get(prefix2 + "forward_snat").(string)
		obj1_3.DropResponseCode = d.Get(prefix2 + "drop_response_code").(int)
		obj1_3.Action1 = d.Get(prefix2 + "action1").(string)
		obj1_3.FakeSg = d.Get(prefix2 + "fake_sg").(string)
		obj1_3.UserTag = d.Get(prefix2 + "user_tag").(string)
		obj1_3.RealSg = d.Get(prefix2 + "real_sg").(string)
		obj1_3.DropMessage = d.Get(prefix2 + "drop_message").(string)

		SamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
		obj1_3.Counters1 = make([]go_thunder.SamplingEnable3, 0, SamplingEnableCount)

		for j := 0; j < SamplingEnableCount; j++ {
			var obj1_3_1 go_thunder.SamplingEnable3
			prefix3 := fmt.Sprintf(prefix2+"sampling_enable.%d.", j)
			obj1_3_1.Counters1 = d.Get(prefix3 + "counters1").(string)
			obj1_3.Counters1 = append(obj1_3.Counters1, obj1_3_1)
		}

		obj1_3.FallBack = d.Get(prefix2 + "fall_back").(string)
		obj1_3.FallBackSnat = d.Get(prefix2 + "fall_back_snat").(string)
		obj1_3.DropRedirectURL = d.Get(prefix2 + "drop_redirect_url").(string)
		obj1_3.Name = d.Get(prefix2 + "name").(string)
		obj1.Log = append(obj1.Log, obj1_3)
	}

	obj1.NoClientConnReuse = d.Get(prefix1 + "no_client_conn_reuse").(int)
	obj1.RequireWebCategory = d.Get(prefix1 + "require_web_category").(int)
	obj1.AcosEventLog = d.Get(prefix1 + "acos_event_log").(int)

	SourceListCount := d.Get(prefix1 + "source_list.#").(int)
	obj1.MatchAny = make([]go_thunder.SourceList, 0, SourceListCount)

	for i := 0; i < SourceListCount; i++ {
		var obj1_4 go_thunder.SourceList
		prefix2 := fmt.Sprintf("forward_policy.0.source_list.%d.", i)
		obj1_4.MatchAny = d.Get(prefix2 + "match_any").(int)
		obj1_4.Name = d.Get(prefix2 + "name").(string)
		obj1_4.MatchAuthorizePolicy = d.Get(prefix2 + "match_authorize_policy").(string)

		var obj1_4_1 go_thunder.Destination
		prefix3 := prefix2 + "destination.0."

		ClassListListCount := d.Get(prefix3 + "class_list_list.#").(int)
		obj1_4_1.UUID = make([]go_thunder.ClassListList, 0, ClassListListCount)

		for j := 0; j < ClassListListCount; j++ {
			var obj1_4_1_1 go_thunder.ClassListList
			prefix4 := fmt.Sprintf(prefix3+"class_list_list.%d.", j)
			obj1_4_1_1.DestClassList = d.Get(prefix4 + "dest_class_list").(string)
			obj1_4_1_1.Priority = d.Get(prefix4 + "priority").(int)

			SamplingEnableCount := d.Get(prefix4 + "sampling_enable.#").(int)
			obj1_4_1_1.Counters1 = make([]go_thunder.SamplingEnable3, 0, SamplingEnableCount)

			for k := 0; k < SamplingEnableCount; k++ {
				var obj1_4_1_1_1 go_thunder.SamplingEnable3
				prefix5 := fmt.Sprintf(prefix4+"sampling_enable.%d.", k)
				obj1_4_1_1_1.Counters1 = d.Get(prefix5 + "counters1").(string)
				obj1_4_1_1.Counters1 = append(obj1_4_1_1.Counters1, obj1_4_1_1_1)
			}

			obj1_4_1_1.Action = d.Get(prefix4 + "action").(string)
			obj1_4_1_1.Type = d.Get(prefix4 + "type").(string)
			obj1_4_1.UUID = append(obj1_4_1.UUID, obj1_4_1_1)
		}

		WebCategoryListListCount := d.Get(prefix3 + "web_category_list_list.#").(int)
		obj1_4_1.WebCategoryList = make([]go_thunder.WebCategoryListList, 0, WebCategoryListListCount)

		for j := 0; j < WebCategoryListListCount; j++ {
			var obj1_4_1_2 go_thunder.WebCategoryListList
			prefix4 := fmt.Sprintf(prefix3+"web_category_list_list.%d.", i)
			obj1_4_1_2.WebCategoryList = d.Get(prefix4 + "web_category_list").(string)
			obj1_4_1_2.Priority = d.Get(prefix4 + "priority").(int)

			SamplingEnableCount := d.Get(prefix4 + "sampling_enable.#").(int)
			obj1_4_1_2.Counters1 = make([]go_thunder.SamplingEnable3, 0, SamplingEnableCount)

			for k := 0; k < SamplingEnableCount; k++ {
				var obj1_4_1_2_1 go_thunder.SamplingEnable3
				prefix5 := fmt.Sprintf(prefix4+"sampling_enable.%d.", k)
				obj1_4_1_2_1.Counters1 = d.Get(prefix5 + "counters1").(string)
				obj1_4_1_2.Counters1 = append(obj1_4_1_2.Counters1, obj1_4_1_2_1)
			}

			obj1_4_1_2.Action = d.Get(prefix4 + "action").(string)
			obj1_4_1_2.Type = d.Get(prefix4 + "type").(string)
			obj1_4_1.WebCategoryList = append(obj1_4_1.WebCategoryList, obj1_4_1_2)
		}

		var obj1_4_1_3 go_thunder.Any
		prefix4 := prefix3 + "any.0."
		obj1_4_1_3.Action = d.Get(prefix4 + "action").(string)

		SamplingEnableCount := d.Get(prefix4 + "sampling_enable.#").(int)
		obj1_4_1_3.Counters1 = make([]go_thunder.SamplingEnable3, 0, SamplingEnableCount)

		for j := 0; j < SamplingEnableCount; j++ {
			var obj1_4_1_3_1 go_thunder.SamplingEnable3
			prefix5 := fmt.Sprintf(prefix4+"sampling_enable.%d.", j)
			obj1_4_1_3_1.Counters1 = d.Get(prefix5 + "counters1").(string)
			obj1_4_1_3.Counters1 = append(obj1_4_1_3.Counters1, obj1_4_1_3_1)
		}

		obj1_4_1.Action = obj1_4_1_3

		obj1_4.WebCategoryList = obj1_4_1

		obj1_4.UserTag = d.Get(prefix2 + "user_tag").(string)
		obj1_4.Priority = d.Get(prefix2 + "priority").(int)

		SamplingEnableCount1 := d.Get(prefix2 + "sampling_enable.#").(int)
		obj1_4.Counters1 = make([]go_thunder.SamplingEnable3, 0, SamplingEnableCount1)

		for j := 0; j < SamplingEnableCount1; j++ {
			var obj1_4_2 go_thunder.SamplingEnable3
			prefix3 := fmt.Sprintf(prefix2+"sampling_enable.%d.", j)
			obj1_4_2.Counters1 = d.Get(prefix3 + "counters1").(string)
			obj1_4.Counters1 = append(obj1_4.Counters1, obj1_4_2)
		}

		obj1_4.MatchClassList = d.Get(prefix2 + "match_class_list").(string)
		obj1.MatchAny = append(obj1.MatchAny, obj1_4)
	}

	c.SsliURLFiltering = obj1

	c.UseDestinationIP = d.Get("use_destination_ip").(int)
	c.Name = d.Get("name").(string)
	c.OverLimit = d.Get("over_limit").(int)

	var obj2 go_thunder.ClassList2
	prefix1 = "class_list.0."
	obj2.HeaderName = d.Get(prefix1 + "header_name").(string)

	LidListCount := d.Get(prefix1 + "lid_list.#").(int)
	obj2.RequestRateLimit = make([]go_thunder.LidList2, 0, LidListCount)

	for i := 0; i < LidListCount; i++ {
		var obj2_1 go_thunder.LidList2
		prefix2 := fmt.Sprintf(prefix1+"lid_list.%d.", i)
		obj2_1.RequestRateLimit = d.Get(prefix2 + "request_rate_limit").(int)
		obj2_1.ActionValue = d.Get(prefix2 + "action_value").(string)
		obj2_1.RequestPer = d.Get(prefix2 + "request_per").(int)
		obj2_1.BwRateLimit = d.Get(prefix2 + "bw_rate_limit").(int)
		obj2_1.ConnLimit = d.Get(prefix2 + "conn_limit").(int)
		obj2_1.Log = d.Get(prefix2 + "log").(int)
		obj2_1.DirectActionValue = d.Get(prefix2 + "direct_action_value").(string)
		obj2_1.ConnPer = d.Get(prefix2 + "conn_per").(int)
		obj2_1.DirectFail = d.Get(prefix2 + "direct_fail").(int)
		obj2_1.ConnRateLimit = d.Get(prefix2 + "conn_rate_limit").(int)
		obj2_1.DirectPbslbLogging = d.Get(prefix2 + "direct_pbslb_logging").(int)

		var obj2_1_1 go_thunder.DNS64
		prefix3 := prefix2 + "dns64.0."
		obj2_1_1.Prefix = d.Get(prefix3 + "prefix").(string)
		obj2_1_1.ExclusiveAnswer = d.Get(prefix3 + "exclusive_answer").(int)
		obj2_1_1.Disable = d.Get(prefix3 + "disable").(int)
		obj2_1.Prefix = obj2_1_1

		obj2_1.Lidnum = d.Get(prefix2 + "lidnum").(int)
		obj2_1.OverLimitAction = d.Get(prefix2 + "over_limit_action").(int)

		ResponseCodeRateLimitCount := d.Get(prefix2 + "response_code_rate_limit.#").(int)
		obj2_1.Threshold = make([]go_thunder.ResponseCodeRateLimit, 0, ResponseCodeRateLimitCount)

		for j := 0; j < ResponseCodeRateLimitCount; j++ {
			var obj2_1_2 go_thunder.ResponseCodeRateLimit
			prefix3 := fmt.Sprintf(prefix2+"response_code_rate_limit.%d.", j)
			obj2_1_2.Threshold = d.Get(prefix3 + "threshold").(int)
			obj2_1_2.CodeRangeEnd = d.Get(prefix3 + "code_range_end").(int)
			obj2_1_2.CodeRangeStart = d.Get(prefix3 + "code_range_start").(int)
			obj2_1_2.Period = d.Get(prefix3 + "period").(int)
			obj2_1.Threshold = append(obj2_1.Threshold, obj2_1_2)
		}

		obj2_1.DirectServiceGroup = d.Get(prefix2 + "direct_service_group").(string)
		obj2_1.RequestLimit = d.Get(prefix2 + "request_limit").(int)
		obj2_1.DirectActionInterval = d.Get(prefix2 + "direct_action_interval").(int)
		obj2_1.BwPer = d.Get(prefix2 + "bw_per").(int)
		obj2_1.Interval = d.Get(prefix2 + "interval").(int)
		obj2_1.UserTag = d.Get(prefix2 + "user_tag").(string)
		obj2_1.DirectAction = d.Get(prefix2 + "direct_action").(int)
		obj2_1.Lockout = d.Get(prefix2 + "lockout").(int)
		obj2_1.DirectLoggingDrpRst = d.Get(prefix2 + "direct_logging_drp_rst").(int)
		obj2_1.DirectPbslbInterval = d.Get(prefix2 + "direct_pbslb_interval").(int)
		obj2.RequestRateLimit = append(obj2.RequestRateLimit, obj2_1)
	}

	obj2.Name = d.Get(prefix1 + "name").(string)
	obj2.ClientIPL3Dest = d.Get(prefix1 + "client_ip_l3_dest").(int)
	obj2.ClientIPL7Header = d.Get(prefix1 + "client_ip_l7_header").(int)
	c.HeaderName = obj2

	c.Interval = d.Get("interval").(int)
	c.Share = d.Get("share").(int)
	c.FullDomainTree = d.Get("full_domain_tree").(int)
	c.OverLimitLogging = d.Get("over_limit_logging").(int)
	c.BwListName = d.Get("bw_list_name").(string)
	c.Timeout = d.Get("timeout").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SamplingEnable3, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj3 go_thunder.SamplingEnable3
		prefix1 := fmt.Sprintf("sampling_enable.%d.", i)
		obj3.Counters1 = d.Get(prefix1 + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj3)
	}

	c.UserTag = d.Get("user_tag").(string)

	BwListIDCount := d.Get("bw_list_id.#").(int)
	c.PbslbInterval = make([]go_thunder.BwListID, 0, BwListIDCount)

	for i := 0; i < BwListIDCount; i++ {
		var obj4 go_thunder.BwListID
		prefix1 := fmt.Sprintf("bw_list_id.%d.", i)
		obj4.PbslbInterval = d.Get(prefix1 + "pbslb_interval").(int)
		obj4.ActionInterval = d.Get(prefix1 + "action_interval").(int)
		obj4.ServiceGroup = d.Get(prefix1 + "service_group").(string)
		obj4.LoggingDrpRst = d.Get(prefix1 + "logging_drp_rst").(int)
		obj4.Fail = d.Get(prefix1 + "fail").(int)
		obj4.PbslbLogging = d.Get(prefix1 + "pbslb_logging").(int)
		obj4.ID = d.Get(prefix1 + "id").(int)
		obj4.BwListAction = d.Get(prefix1 + "bw_list_action").(string)
		c.PbslbInterval = append(c.PbslbInterval, obj4)
	}

	c.OverLimitLockup = d.Get("over_limit_lockup").(int)
	c.OverLimitReset = d.Get("over_limit_reset").(int)
	c.Overlap = d.Get("overlap").(int)

	vc.UUID = c
	return vc
}
