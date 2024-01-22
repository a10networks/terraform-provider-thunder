package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_oper`: Operational Status for the object web-category\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryOperRead,

		Schema: map[string]*schema.Schema{
			"bypassed_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"number_of_urls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"all_urls": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"url_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"intercepted_urls": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"number_of_urls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"all_urls": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"url_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"license": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"module_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"license_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"license_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"license_expiry": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remaining_period": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_grace": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"grace_period": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"serial_number": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"web_cat_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_database_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_database_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_database_size": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_database_version": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"web_cat_last_update_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_next_update_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_connection_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_failure_reason": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"web_cat_last_successful_connection": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"statistics": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_dplane_threads": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"num_lookup_threads": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"per_cpu_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"req_queue": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"req_dropped": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"req_processed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"req_lookup_processed": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"total_req_queue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_req_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_req_processed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_req_lookup_processed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"clear_cache": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"url": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"category": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_db_only": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"web_reputation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"number_of_urls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"all_urls": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"url_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"intercepted_urls": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"url_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"number_of_urls": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"all_urls": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"url_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"bypassed_urls": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"url_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"number_of_urls": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"all_urls": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"url_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"url": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"reputation_score": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_db_only": {
													Type: schema.TypeInt, Optional: true, Description: "",
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
		},
	}
}

func resourceWebCategoryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryOperBypassedUrls := setObjectWebCategoryOperBypassedUrls(res)
		d.Set("bypassed_urls", WebCategoryOperBypassedUrls)
		WebCategoryOperInterceptedUrls := setObjectWebCategoryOperInterceptedUrls(res)
		d.Set("intercepted_urls", WebCategoryOperInterceptedUrls)
		WebCategoryOperLicense := setObjectWebCategoryOperLicense(res)
		d.Set("license", WebCategoryOperLicense)
		WebCategoryOperOper := setObjectWebCategoryOperOper(res)
		d.Set("oper", WebCategoryOperOper)
		WebCategoryOperStatistics := setObjectWebCategoryOperStatistics(res)
		d.Set("statistics", WebCategoryOperStatistics)
		WebCategoryOperUrl := setObjectWebCategoryOperUrl(res)
		d.Set("url", WebCategoryOperUrl)
		WebCategoryOperWebReputation := setObjectWebCategoryOperWebReputation(res)
		d.Set("web_reputation", WebCategoryOperWebReputation)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryOperBypassedUrls(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryOperBypassedUrlsOper(ret.DtWebCategoryOper.BypassedUrls.Oper),
		},
	}
}

func setObjectWebCategoryOperBypassedUrlsOper(d edpt.WebCategoryOperBypassedUrlsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryOperBypassedUrlsOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperBypassedUrlsOperUrlList(d []edpt.WebCategoryOperBypassedUrlsOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperInterceptedUrls(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryOperInterceptedUrlsOper(ret.DtWebCategoryOper.InterceptedUrls.Oper),
		},
	}
}

func setObjectWebCategoryOperInterceptedUrlsOper(d edpt.WebCategoryOperInterceptedUrlsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryOperInterceptedUrlsOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperInterceptedUrlsOperUrlList(d []edpt.WebCategoryOperInterceptedUrlsOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperLicense(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryOperLicenseOper(ret.DtWebCategoryOper.License.Oper),
		},
	}
}

func setObjectWebCategoryOperLicenseOper(d edpt.WebCategoryOperLicenseOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["module_status"] = d.ModuleStatus

	in["license_status"] = d.LicenseStatus

	in["license_type"] = d.LicenseType

	in["license_expiry"] = d.LicenseExpiry

	in["remaining_period"] = d.RemainingPeriod

	in["is_grace"] = d.IsGrace

	in["grace_period"] = d.GracePeriod

	in["serial_number"] = d.SerialNumber
	result = append(result, in)
	return result
}

func setObjectWebCategoryOperOper(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"web_cat_version":                    ret.DtWebCategoryOper.Oper.WebCatVersion,
			"web_cat_database_name":              ret.DtWebCategoryOper.Oper.WebCatDatabaseName,
			"web_cat_database_status":            ret.DtWebCategoryOper.Oper.WebCatDatabaseStatus,
			"web_cat_database_size":              ret.DtWebCategoryOper.Oper.WebCatDatabaseSize,
			"web_cat_database_version":           ret.DtWebCategoryOper.Oper.WebCatDatabaseVersion,
			"web_cat_last_update_time":           ret.DtWebCategoryOper.Oper.WebCatLastUpdateTime,
			"web_cat_next_update_time":           ret.DtWebCategoryOper.Oper.WebCatNextUpdateTime,
			"web_cat_connection_status":          ret.DtWebCategoryOper.Oper.WebCatConnectionStatus,
			"web_cat_failure_reason":             ret.DtWebCategoryOper.Oper.WebCatFailureReason,
			"web_cat_last_successful_connection": ret.DtWebCategoryOper.Oper.WebCatLastSuccessfulConnection,
		},
	}
}

func setObjectWebCategoryOperStatistics(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryOperStatisticsOper(ret.DtWebCategoryOper.Statistics.Oper),
		},
	}
}

func setObjectWebCategoryOperStatisticsOper(d edpt.WebCategoryOperStatisticsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["num_dplane_threads"] = d.NumDplaneThreads

	in["num_lookup_threads"] = d.NumLookupThreads
	in["per_cpu_list"] = setSliceWebCategoryOperStatisticsOperPerCpuList(d.PerCpuList)

	in["total_req_queue"] = d.TotalReqQueue

	in["total_req_dropped"] = d.TotalReqDropped

	in["total_req_processed"] = d.TotalReqProcessed

	in["total_req_lookup_processed"] = d.TotalReqLookupProcessed

	in["clear_cache"] = d.ClearCache
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperStatisticsOperPerCpuList(d []edpt.WebCategoryOperStatisticsOperPerCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["req_queue"] = item.ReqQueue
		in["req_dropped"] = item.ReqDropped
		in["req_processed"] = item.ReqProcessed
		in["req_lookup_processed"] = item.ReqLookupProcessed
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperUrl(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectWebCategoryOperUrlOper(ret.DtWebCategoryOper.Url.Oper),
		},
	}
}

func setObjectWebCategoryOperUrlOper(d edpt.WebCategoryOperUrlOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["category_list"] = setSliceWebCategoryOperUrlOperCategoryList(d.CategoryList)

	in["name"] = d.Name

	in["local_db_only"] = d.LocalDbOnly
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperUrlOperCategoryList(d []edpt.WebCategoryOperUrlOperCategoryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["category"] = item.Category
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperWebReputation(ret edpt.DataWebCategoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":             setObjectWebCategoryOperWebReputationOper(ret.DtWebCategoryOper.WebReputation.Oper),
			"intercepted_urls": setObjectWebCategoryOperWebReputationInterceptedUrls(ret.DtWebCategoryOper.WebReputation.InterceptedUrls),
			"bypassed_urls":    setObjectWebCategoryOperWebReputationBypassedUrls(ret.DtWebCategoryOper.WebReputation.BypassedUrls),
			"url":              setObjectWebCategoryOperWebReputationUrl(ret.DtWebCategoryOper.WebReputation.Url),
		},
	}
}

func setObjectWebCategoryOperWebReputationOper(d edpt.WebCategoryOperWebReputationOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryOperWebReputationOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperWebReputationOperUrlList(d []edpt.WebCategoryOperWebReputationOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperWebReputationInterceptedUrls(d edpt.WebCategoryOperWebReputationInterceptedUrls) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectWebCategoryOperWebReputationInterceptedUrlsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectWebCategoryOperWebReputationInterceptedUrlsOper(d edpt.WebCategoryOperWebReputationInterceptedUrlsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryOperWebReputationInterceptedUrlsOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperWebReputationInterceptedUrlsOperUrlList(d []edpt.WebCategoryOperWebReputationInterceptedUrlsOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperWebReputationBypassedUrls(d edpt.WebCategoryOperWebReputationBypassedUrls) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectWebCategoryOperWebReputationBypassedUrlsOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectWebCategoryOperWebReputationBypassedUrlsOper(d edpt.WebCategoryOperWebReputationBypassedUrlsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["url_list"] = setSliceWebCategoryOperWebReputationBypassedUrlsOperUrlList(d.UrlList)

	in["number_of_urls"] = d.NumberOfUrls

	in["all_urls"] = d.AllUrls

	in["url_name"] = d.UrlName
	result = append(result, in)
	return result
}

func setSliceWebCategoryOperWebReputationBypassedUrlsOperUrlList(d []edpt.WebCategoryOperWebReputationBypassedUrlsOperUrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["url_name"] = item.UrlName
		result = append(result, in)
	}
	return result
}

func setObjectWebCategoryOperWebReputationUrl(d edpt.WebCategoryOperWebReputationUrl) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectWebCategoryOperWebReputationUrlOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectWebCategoryOperWebReputationUrlOper(d edpt.WebCategoryOperWebReputationUrlOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["reputation_score"] = d.ReputationScore

	in["name"] = d.Name

	in["local_db_only"] = d.LocalDbOnly
	result = append(result, in)
	return result
}

func getObjectWebCategoryOperBypassedUrls(d []interface{}) edpt.WebCategoryOperBypassedUrls {

	count1 := len(d)
	var ret edpt.WebCategoryOperBypassedUrls
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperBypassedUrlsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperBypassedUrlsOper(d []interface{}) edpt.WebCategoryOperBypassedUrlsOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperBypassedUrlsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryOperBypassedUrlsOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryOperBypassedUrlsOperUrlList(d []interface{}) []edpt.WebCategoryOperBypassedUrlsOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperBypassedUrlsOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperBypassedUrlsOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperInterceptedUrls(d []interface{}) edpt.WebCategoryOperInterceptedUrls {

	count1 := len(d)
	var ret edpt.WebCategoryOperInterceptedUrls
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperInterceptedUrlsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperInterceptedUrlsOper(d []interface{}) edpt.WebCategoryOperInterceptedUrlsOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperInterceptedUrlsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryOperInterceptedUrlsOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryOperInterceptedUrlsOperUrlList(d []interface{}) []edpt.WebCategoryOperInterceptedUrlsOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperInterceptedUrlsOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperInterceptedUrlsOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperLicense(d []interface{}) edpt.WebCategoryOperLicense {

	count1 := len(d)
	var ret edpt.WebCategoryOperLicense
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperLicenseOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperLicenseOper(d []interface{}) edpt.WebCategoryOperLicenseOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperLicenseOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ModuleStatus = in["module_status"].(string)
		ret.LicenseStatus = in["license_status"].(string)
		ret.LicenseType = in["license_type"].(string)
		ret.LicenseExpiry = in["license_expiry"].(string)
		ret.RemainingPeriod = in["remaining_period"].(string)
		ret.IsGrace = in["is_grace"].(string)
		ret.GracePeriod = in["grace_period"].(string)
		ret.SerialNumber = in["serial_number"].(string)
	}
	return ret
}

func getObjectWebCategoryOperOper(d []interface{}) edpt.WebCategoryOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WebCatVersion = in["web_cat_version"].(string)
		ret.WebCatDatabaseName = in["web_cat_database_name"].(string)
		ret.WebCatDatabaseStatus = in["web_cat_database_status"].(string)
		ret.WebCatDatabaseSize = in["web_cat_database_size"].(string)
		ret.WebCatDatabaseVersion = in["web_cat_database_version"].(int)
		ret.WebCatLastUpdateTime = in["web_cat_last_update_time"].(string)
		ret.WebCatNextUpdateTime = in["web_cat_next_update_time"].(string)
		ret.WebCatConnectionStatus = in["web_cat_connection_status"].(string)
		ret.WebCatFailureReason = in["web_cat_failure_reason"].(string)
		ret.WebCatLastSuccessfulConnection = in["web_cat_last_successful_connection"].(string)
	}
	return ret
}

func getObjectWebCategoryOperStatistics(d []interface{}) edpt.WebCategoryOperStatistics {

	count1 := len(d)
	var ret edpt.WebCategoryOperStatistics
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperStatisticsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperStatisticsOper(d []interface{}) edpt.WebCategoryOperStatisticsOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperStatisticsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumDplaneThreads = in["num_dplane_threads"].(int)
		ret.NumLookupThreads = in["num_lookup_threads"].(int)
		ret.PerCpuList = getSliceWebCategoryOperStatisticsOperPerCpuList(in["per_cpu_list"].([]interface{}))
		ret.TotalReqQueue = in["total_req_queue"].(int)
		ret.TotalReqDropped = in["total_req_dropped"].(int)
		ret.TotalReqProcessed = in["total_req_processed"].(int)
		ret.TotalReqLookupProcessed = in["total_req_lookup_processed"].(int)
		ret.ClearCache = in["clear_cache"].(string)
	}
	return ret
}

func getSliceWebCategoryOperStatisticsOperPerCpuList(d []interface{}) []edpt.WebCategoryOperStatisticsOperPerCpuList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperStatisticsOperPerCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperStatisticsOperPerCpuList
		oi.ReqQueue = in["req_queue"].(int)
		oi.ReqDropped = in["req_dropped"].(int)
		oi.ReqProcessed = in["req_processed"].(int)
		oi.ReqLookupProcessed = in["req_lookup_processed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperUrl(d []interface{}) edpt.WebCategoryOperUrl {

	count1 := len(d)
	var ret edpt.WebCategoryOperUrl
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperUrlOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperUrlOper(d []interface{}) edpt.WebCategoryOperUrlOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperUrlOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CategoryList = getSliceWebCategoryOperUrlOperCategoryList(in["category_list"].([]interface{}))
		ret.Name = in["name"].(string)
		ret.LocalDbOnly = in["local_db_only"].(int)
	}
	return ret
}

func getSliceWebCategoryOperUrlOperCategoryList(d []interface{}) []edpt.WebCategoryOperUrlOperCategoryList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperUrlOperCategoryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperUrlOperCategoryList
		oi.Category = in["category"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperWebReputation(d []interface{}) edpt.WebCategoryOperWebReputation {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperWebReputationOper(in["oper"].([]interface{}))
		ret.InterceptedUrls = getObjectWebCategoryOperWebReputationInterceptedUrls(in["intercepted_urls"].([]interface{}))
		ret.BypassedUrls = getObjectWebCategoryOperWebReputationBypassedUrls(in["bypassed_urls"].([]interface{}))
		ret.Url = getObjectWebCategoryOperWebReputationUrl(in["url"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperWebReputationOper(d []interface{}) edpt.WebCategoryOperWebReputationOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryOperWebReputationOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryOperWebReputationOperUrlList(d []interface{}) []edpt.WebCategoryOperWebReputationOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperWebReputationOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperWebReputationOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperWebReputationInterceptedUrls(d []interface{}) edpt.WebCategoryOperWebReputationInterceptedUrls {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationInterceptedUrls
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperWebReputationInterceptedUrlsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperWebReputationInterceptedUrlsOper(d []interface{}) edpt.WebCategoryOperWebReputationInterceptedUrlsOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationInterceptedUrlsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryOperWebReputationInterceptedUrlsOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryOperWebReputationInterceptedUrlsOperUrlList(d []interface{}) []edpt.WebCategoryOperWebReputationInterceptedUrlsOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperWebReputationInterceptedUrlsOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperWebReputationInterceptedUrlsOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperWebReputationBypassedUrls(d []interface{}) edpt.WebCategoryOperWebReputationBypassedUrls {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationBypassedUrls
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperWebReputationBypassedUrlsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperWebReputationBypassedUrlsOper(d []interface{}) edpt.WebCategoryOperWebReputationBypassedUrlsOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationBypassedUrlsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UrlList = getSliceWebCategoryOperWebReputationBypassedUrlsOperUrlList(in["url_list"].([]interface{}))
		ret.NumberOfUrls = in["number_of_urls"].(int)
		ret.AllUrls = in["all_urls"].(string)
		ret.UrlName = in["url_name"].(string)
	}
	return ret
}

func getSliceWebCategoryOperWebReputationBypassedUrlsOperUrlList(d []interface{}) []edpt.WebCategoryOperWebReputationBypassedUrlsOperUrlList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryOperWebReputationBypassedUrlsOperUrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryOperWebReputationBypassedUrlsOperUrlList
		oi.UrlName = in["url_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebCategoryOperWebReputationUrl(d []interface{}) edpt.WebCategoryOperWebReputationUrl {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationUrl
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectWebCategoryOperWebReputationUrlOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectWebCategoryOperWebReputationUrlOper(d []interface{}) edpt.WebCategoryOperWebReputationUrlOper {

	count1 := len(d)
	var ret edpt.WebCategoryOperWebReputationUrlOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ReputationScore = in["reputation_score"].(string)
		ret.Name = in["name"].(string)
		ret.LocalDbOnly = in["local_db_only"].(int)
	}
	return ret
}

func dataToEndpointWebCategoryOper(d *schema.ResourceData) edpt.WebCategoryOper {
	var ret edpt.WebCategoryOper

	ret.BypassedUrls = getObjectWebCategoryOperBypassedUrls(d.Get("bypassed_urls").([]interface{}))

	ret.InterceptedUrls = getObjectWebCategoryOperInterceptedUrls(d.Get("intercepted_urls").([]interface{}))

	ret.License = getObjectWebCategoryOperLicense(d.Get("license").([]interface{}))

	ret.Oper = getObjectWebCategoryOperOper(d.Get("oper").([]interface{}))

	ret.Statistics = getObjectWebCategoryOperStatistics(d.Get("statistics").([]interface{}))

	ret.Url = getObjectWebCategoryOperUrl(d.Get("url").([]interface{}))

	ret.WebReputation = getObjectWebCategoryOperWebReputation(d.Get("web_reputation").([]interface{}))
	return ret
}
