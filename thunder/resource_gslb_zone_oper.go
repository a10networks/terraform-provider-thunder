package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_oper`: Operational Status for the object zone\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneOperRead,

		Schema: map[string]*schema.Schema{
			"dns_caa_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical_flag": {
							Type: schema.TypeInt, Required: true, Description: "Issuer Critical Flag",
						},
						"property_tag": {
							Type: schema.TypeString, Required: true, Description: "Specify other property tags, only allowed lowercase alphanumeric",
						},
						"rdata": {
							Type: schema.TypeString, Required: true, Description: "Specify the Issuer Domain Name or a URL",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_server": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"dns_mx_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mx_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_server": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"dns_ns_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ns_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_server": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify the name for the DNS zone",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dns_soa_record_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"expire": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"refresh": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serial": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mx_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"service_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_port": {
							Type: schema.TypeInt, Required: true, Description: "Port number of the service",
						},
						"service_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the service name for the zone, * for wildcard",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cache_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alias": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"cache_length": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cache_ttl": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cache_dns_flag": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"question_records": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"answer_records": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"authority_records": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"additional_records": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"session_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"best": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"last_second_hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ttl": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"update": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"aging": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"matched": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_sessions": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dns_a_record_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rec_ttl": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"dns_mx_record_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mx_name": {
										Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
									},
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"last_server": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hits": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"dns_ns_record_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ns_name": {
										Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
									},
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"last_server": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hits": {
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

func resourceGslbZoneOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneOperDnsCaaRecordList := setSliceGslbZoneOperDnsCaaRecordList(res)
		d.Set("dns_caa_record_list", GslbZoneOperDnsCaaRecordList)
		GslbZoneOperDnsMxRecordList := setSliceGslbZoneOperDnsMxRecordList(res)
		d.Set("dns_mx_record_list", GslbZoneOperDnsMxRecordList)
		GslbZoneOperDnsNsRecordList := setSliceGslbZoneOperDnsNsRecordList(res)
		d.Set("dns_ns_record_list", GslbZoneOperDnsNsRecordList)
		GslbZoneOperOper := setObjectGslbZoneOperOper(res)
		d.Set("oper", GslbZoneOperOper)
		GslbZoneOperServiceList := setSliceGslbZoneOperServiceList(res)
		d.Set("service_list", GslbZoneOperServiceList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbZoneOperDnsCaaRecordList(d edpt.DataGslbZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneOper.DnsCaaRecordList {
		in := make(map[string]interface{})
		in["critical_flag"] = item.CriticalFlag
		in["property_tag"] = item.PropertyTag
		in["rdata"] = item.Rdata
		in["oper"] = setObjectGslbZoneOperDnsCaaRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneOperDnsCaaRecordListOper(d edpt.GslbZoneOperDnsCaaRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer
	result = append(result, in)
	return result
}

func setSliceGslbZoneOperDnsMxRecordList(d edpt.DataGslbZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneOper.DnsMxRecordList {
		in := make(map[string]interface{})
		in["mx_name"] = item.MxName
		in["oper"] = setObjectGslbZoneOperDnsMxRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneOperDnsMxRecordListOper(d edpt.GslbZoneOperDnsMxRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer

	in["hits"] = d.Hits

	in["priority"] = d.Priority
	result = append(result, in)
	return result
}

func setSliceGslbZoneOperDnsNsRecordList(d edpt.DataGslbZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneOper.DnsNsRecordList {
		in := make(map[string]interface{})
		in["ns_name"] = item.NsName
		in["oper"] = setObjectGslbZoneOperDnsNsRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneOperDnsNsRecordListOper(d edpt.GslbZoneOperDnsNsRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setObjectGslbZoneOperOper(ret edpt.DataGslbZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":               ret.DtGslbZoneOper.Oper.State,
			"dns_soa_record_list": setSliceGslbZoneOperOperDnsSoaRecordList(ret.DtGslbZoneOper.Oper.DnsSoaRecordList),
		},
	}
}

func setSliceGslbZoneOperOperDnsSoaRecordList(d []edpt.GslbZoneOperOperDnsSoaRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["type"] = item.Type
		in["expire"] = item.Expire
		in["refresh"] = item.Refresh
		in["serial"] = item.Serial
		in["retry"] = item.Retry
		in["ttl"] = item.Ttl
		in["mx_name"] = item.MxName
		result = append(result, in)
	}
	return result
}

func setSliceGslbZoneOperServiceList(d edpt.DataGslbZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneOper.ServiceList {
		in := make(map[string]interface{})
		in["service_port"] = item.ServicePort
		in["service_name"] = item.ServiceName
		in["oper"] = setObjectGslbZoneOperServiceListOper(item.Oper)
		in["dns_mx_record_list"] = setSliceGslbZoneOperServiceListDnsMxRecordList(item.DnsMxRecordList)
		in["dns_ns_record_list"] = setSliceGslbZoneOperServiceListDnsNsRecordList(item.DnsNsRecordList)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneOperServiceListOper(d edpt.GslbZoneOperServiceListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State
	in["cache_list"] = setSliceGslbZoneOperServiceListOperCacheList(d.CacheList)
	in["session_list"] = setSliceGslbZoneOperServiceListOperSessionList(d.SessionList)

	in["matched"] = d.Matched

	in["total_sessions"] = d.TotalSessions
	in["dns_a_record_list"] = setSliceGslbZoneOperServiceListOperDnsARecordList(d.DnsARecordList)
	result = append(result, in)
	return result
}

func setSliceGslbZoneOperServiceListOperCacheList(d []edpt.GslbZoneOperServiceListOperCacheList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["alias"] = item.Alias
		in["cache_length"] = item.CacheLength
		in["cache_ttl"] = item.CacheTtl
		in["cache_dns_flag"] = item.CacheDnsFlag
		in["question_records"] = item.QuestionRecords
		in["answer_records"] = item.AnswerRecords
		in["authority_records"] = item.AuthorityRecords
		in["additional_records"] = item.AdditionalRecords
		result = append(result, in)
	}
	return result
}

func setSliceGslbZoneOperServiceListOperSessionList(d []edpt.GslbZoneOperServiceListOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["client"] = item.Client
		in["best"] = item.Best
		in["mode"] = item.Mode
		in["hits"] = item.Hits
		in["last_second_hits"] = item.LastSecondHits
		in["ttl"] = item.Ttl
		in["update"] = item.Update
		in["aging"] = item.Aging
		result = append(result, in)
	}
	return result
}

func setSliceGslbZoneOperServiceListOperDnsARecordList(d []edpt.GslbZoneOperServiceListOperDnsARecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		in["rec_ttl"] = item.RecTtl
		result = append(result, in)
	}
	return result
}

func setSliceGslbZoneOperServiceListDnsMxRecordList(d []edpt.GslbZoneOperServiceListDnsMxRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["mx_name"] = item.MxName
		in["oper"] = setObjectGslbZoneOperServiceListDnsMxRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneOperServiceListDnsMxRecordListOper(d edpt.GslbZoneOperServiceListDnsMxRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer

	in["hits"] = d.Hits

	in["priority"] = d.Priority
	result = append(result, in)
	return result
}

func setSliceGslbZoneOperServiceListDnsNsRecordList(d []edpt.GslbZoneOperServiceListDnsNsRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ns_name"] = item.NsName
		in["oper"] = setObjectGslbZoneOperServiceListDnsNsRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneOperServiceListDnsNsRecordListOper(d edpt.GslbZoneOperServiceListDnsNsRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func getSliceGslbZoneOperDnsCaaRecordList(d []interface{}) []edpt.GslbZoneOperDnsCaaRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperDnsCaaRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperDnsCaaRecordList
		oi.CriticalFlag = in["critical_flag"].(int)
		oi.PropertyTag = in["property_tag"].(string)
		oi.Rdata = in["rdata"].(string)
		oi.Oper = getObjectGslbZoneOperDnsCaaRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneOperDnsCaaRecordListOper(d []interface{}) edpt.GslbZoneOperDnsCaaRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperDnsCaaRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
	}
	return ret
}

func getSliceGslbZoneOperDnsMxRecordList(d []interface{}) []edpt.GslbZoneOperDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Oper = getObjectGslbZoneOperDnsMxRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneOperDnsMxRecordListOper(d []interface{}) edpt.GslbZoneOperDnsMxRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperDnsMxRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
		ret.Priority = in["priority"].(int)
	}
	return ret
}

func getSliceGslbZoneOperDnsNsRecordList(d []interface{}) []edpt.GslbZoneOperDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Oper = getObjectGslbZoneOperDnsNsRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneOperDnsNsRecordListOper(d []interface{}) edpt.GslbZoneOperDnsNsRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperDnsNsRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getObjectGslbZoneOperOper(d []interface{}) edpt.GslbZoneOperOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.DnsSoaRecordList = getSliceGslbZoneOperOperDnsSoaRecordList(in["dns_soa_record_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbZoneOperOperDnsSoaRecordList(d []interface{}) []edpt.GslbZoneOperOperDnsSoaRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperOperDnsSoaRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperOperDnsSoaRecordList
		oi.Name = in["name"].(string)
		oi.Type = in["type"].(string)
		oi.Expire = in["expire"].(int)
		oi.Refresh = in["refresh"].(int)
		oi.Serial = in["serial"].(int)
		oi.Retry = in["retry"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.MxName = in["mx_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneOperServiceList(d []interface{}) []edpt.GslbZoneOperServiceList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperServiceList
		oi.ServicePort = in["service_port"].(int)
		oi.ServiceName = in["service_name"].(string)
		oi.Oper = getObjectGslbZoneOperServiceListOper(in["oper"].([]interface{}))
		oi.DnsMxRecordList = getSliceGslbZoneOperServiceListDnsMxRecordList(in["dns_mx_record_list"].([]interface{}))
		oi.DnsNsRecordList = getSliceGslbZoneOperServiceListDnsNsRecordList(in["dns_ns_record_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneOperServiceListOper(d []interface{}) edpt.GslbZoneOperServiceListOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperServiceListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.CacheList = getSliceGslbZoneOperServiceListOperCacheList(in["cache_list"].([]interface{}))
		ret.SessionList = getSliceGslbZoneOperServiceListOperSessionList(in["session_list"].([]interface{}))
		ret.Matched = in["matched"].(int)
		ret.TotalSessions = in["total_sessions"].(int)
		ret.DnsARecordList = getSliceGslbZoneOperServiceListOperDnsARecordList(in["dns_a_record_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbZoneOperServiceListOperCacheList(d []interface{}) []edpt.GslbZoneOperServiceListOperCacheList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperServiceListOperCacheList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperServiceListOperCacheList
		oi.Alias = in["alias"].(string)
		oi.CacheLength = in["cache_length"].(int)
		oi.CacheTtl = in["cache_ttl"].(int)
		oi.CacheDnsFlag = in["cache_dns_flag"].(string)
		oi.QuestionRecords = in["question_records"].(int)
		oi.AnswerRecords = in["answer_records"].(int)
		oi.AuthorityRecords = in["authority_records"].(int)
		oi.AdditionalRecords = in["additional_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneOperServiceListOperSessionList(d []interface{}) []edpt.GslbZoneOperServiceListOperSessionList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperServiceListOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperServiceListOperSessionList
		oi.Client = in["client"].(string)
		oi.Best = in["best"].(string)
		oi.Mode = in["mode"].(string)
		oi.Hits = in["hits"].(int)
		oi.LastSecondHits = in["last_second_hits"].(int)
		oi.Ttl = in["ttl"].(string)
		oi.Update = in["update"].(int)
		oi.Aging = in["aging"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneOperServiceListOperDnsARecordList(d []interface{}) []edpt.GslbZoneOperServiceListOperDnsARecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperServiceListOperDnsARecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperServiceListOperDnsARecordList
		oi.Ip = in["ip"].(string)
		oi.RecTtl = in["rec_ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbZoneOperServiceListDnsMxRecordList(d []interface{}) []edpt.GslbZoneOperServiceListDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperServiceListDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperServiceListDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Oper = getObjectGslbZoneOperServiceListDnsMxRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneOperServiceListDnsMxRecordListOper(d []interface{}) edpt.GslbZoneOperServiceListDnsMxRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperServiceListDnsMxRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
		ret.Priority = in["priority"].(int)
	}
	return ret
}

func getSliceGslbZoneOperServiceListDnsNsRecordList(d []interface{}) []edpt.GslbZoneOperServiceListDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneOperServiceListDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneOperServiceListDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Oper = getObjectGslbZoneOperServiceListDnsNsRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneOperServiceListDnsNsRecordListOper(d []interface{}) edpt.GslbZoneOperServiceListDnsNsRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneOperServiceListDnsNsRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneOper(d *schema.ResourceData) edpt.GslbZoneOper {
	var ret edpt.GslbZoneOper

	ret.DnsCaaRecordList = getSliceGslbZoneOperDnsCaaRecordList(d.Get("dns_caa_record_list").([]interface{}))

	ret.DnsMxRecordList = getSliceGslbZoneOperDnsMxRecordList(d.Get("dns_mx_record_list").([]interface{}))

	ret.DnsNsRecordList = getSliceGslbZoneOperDnsNsRecordList(d.Get("dns_ns_record_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectGslbZoneOperOper(d.Get("oper").([]interface{}))

	ret.ServiceList = getSliceGslbZoneOperServiceList(d.Get("service_list").([]interface{}))
	return ret
}
