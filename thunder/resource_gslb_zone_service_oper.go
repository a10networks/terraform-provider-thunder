package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_oper`: Operational Status for the object service\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceOperRead,

		Schema: map[string]*schema.Schema{
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
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the service name for the zone, * for wildcard",
			},
			"service_port": {
				Type: schema.TypeInt, Required: true, Description: "Port number of the service",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneServiceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceOperDnsMxRecordList := setSliceGslbZoneServiceOperDnsMxRecordList(res)
		d.Set("dns_mx_record_list", GslbZoneServiceOperDnsMxRecordList)
		GslbZoneServiceOperDnsNsRecordList := setSliceGslbZoneServiceOperDnsNsRecordList(res)
		d.Set("dns_ns_record_list", GslbZoneServiceOperDnsNsRecordList)
		GslbZoneServiceOperOper := setObjectGslbZoneServiceOperOper(res)
		d.Set("oper", GslbZoneServiceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbZoneServiceOperDnsMxRecordList(d edpt.DataGslbZoneServiceOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceOper.DnsMxRecordList {
		in := make(map[string]interface{})
		in["mx_name"] = item.MxName
		in["oper"] = setObjectGslbZoneServiceOperDnsMxRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceOperDnsMxRecordListOper(d edpt.GslbZoneServiceOperDnsMxRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer

	in["hits"] = d.Hits

	in["priority"] = d.Priority
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceOperDnsNsRecordList(d edpt.DataGslbZoneServiceOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceOper.DnsNsRecordList {
		in := make(map[string]interface{})
		in["ns_name"] = item.NsName
		in["oper"] = setObjectGslbZoneServiceOperDnsNsRecordListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceOperDnsNsRecordListOper(d edpt.GslbZoneServiceOperDnsNsRecordListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["last_server"] = d.LastServer

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setObjectGslbZoneServiceOperOper(ret edpt.DataGslbZoneServiceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":             ret.DtGslbZoneServiceOper.Oper.State,
			"cache_list":        setSliceGslbZoneServiceOperOperCacheList(ret.DtGslbZoneServiceOper.Oper.CacheList),
			"session_list":      setSliceGslbZoneServiceOperOperSessionList(ret.DtGslbZoneServiceOper.Oper.SessionList),
			"matched":           ret.DtGslbZoneServiceOper.Oper.Matched,
			"total_sessions":    ret.DtGslbZoneServiceOper.Oper.TotalSessions,
			"dns_a_record_list": setSliceGslbZoneServiceOperOperDnsARecordList(ret.DtGslbZoneServiceOper.Oper.DnsARecordList),
		},
	}
}

func setSliceGslbZoneServiceOperOperCacheList(d []edpt.GslbZoneServiceOperOperCacheList) []map[string]interface{} {
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

func setSliceGslbZoneServiceOperOperSessionList(d []edpt.GslbZoneServiceOperOperSessionList) []map[string]interface{} {
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

func setSliceGslbZoneServiceOperOperDnsARecordList(d []edpt.GslbZoneServiceOperOperDnsARecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip"] = item.Ip
		in["rec_ttl"] = item.RecTtl
		result = append(result, in)
	}
	return result
}

func getSliceGslbZoneServiceOperDnsMxRecordList(d []interface{}) []edpt.GslbZoneServiceOperDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceOperDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceOperDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Oper = getObjectGslbZoneServiceOperDnsMxRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceOperDnsMxRecordListOper(d []interface{}) edpt.GslbZoneServiceOperDnsMxRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneServiceOperDnsMxRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
		ret.Priority = in["priority"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceOperDnsNsRecordList(d []interface{}) []edpt.GslbZoneServiceOperDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceOperDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceOperDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Oper = getObjectGslbZoneServiceOperDnsNsRecordListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceOperDnsNsRecordListOper(d []interface{}) edpt.GslbZoneServiceOperDnsNsRecordListOper {

	count1 := len(d)
	var ret edpt.GslbZoneServiceOperDnsNsRecordListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LastServer = in["last_server"].(string)
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getObjectGslbZoneServiceOperOper(d []interface{}) edpt.GslbZoneServiceOperOper {

	count1 := len(d)
	var ret edpt.GslbZoneServiceOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.CacheList = getSliceGslbZoneServiceOperOperCacheList(in["cache_list"].([]interface{}))
		ret.SessionList = getSliceGslbZoneServiceOperOperSessionList(in["session_list"].([]interface{}))
		ret.Matched = in["matched"].(int)
		ret.TotalSessions = in["total_sessions"].(int)
		ret.DnsARecordList = getSliceGslbZoneServiceOperOperDnsARecordList(in["dns_a_record_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbZoneServiceOperOperCacheList(d []interface{}) []edpt.GslbZoneServiceOperOperCacheList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceOperOperCacheList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceOperOperCacheList
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

func getSliceGslbZoneServiceOperOperSessionList(d []interface{}) []edpt.GslbZoneServiceOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceOperOperSessionList
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

func getSliceGslbZoneServiceOperOperDnsARecordList(d []interface{}) []edpt.GslbZoneServiceOperOperDnsARecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceOperOperDnsARecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceOperOperDnsARecordList
		oi.Ip = in["ip"].(string)
		oi.RecTtl = in["rec_ttl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceOper(d *schema.ResourceData) edpt.GslbZoneServiceOper {
	var ret edpt.GslbZoneServiceOper

	ret.DnsMxRecordList = getSliceGslbZoneServiceOperDnsMxRecordList(d.Get("dns_mx_record_list").([]interface{}))

	ret.DnsNsRecordList = getSliceGslbZoneServiceOperDnsNsRecordList(d.Get("dns_ns_record_list").([]interface{}))

	ret.Oper = getObjectGslbZoneServiceOperOper(d.Get("oper").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(int)

	ret.Name = d.Get("name").(string)
	return ret
}
