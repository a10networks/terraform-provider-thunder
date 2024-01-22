package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsCacheOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dns_cache_oper`: Operational Status for the object dns-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDnsCacheOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache_client": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"unit_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"curr_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"over_rate_limit_times": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lockup": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lockup_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cache_entry": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dnssec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cache_negative": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cache_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cache_class": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"q_length": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"r_length": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ttl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"an_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ns_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ar_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"entry_record": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"record_type": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"record_class": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"record_ttl": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"record_rdlen": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"record_rdata": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"record_rdata_tc": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"client": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entry": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"global": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cache_content": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vport": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vs_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"port_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"port_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"type_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fqdn_domain": {
							Type: schema.TypeString, Optional: true, Description: "domain name",
						},
						"class_string": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"class_value": {
							Type: schema.TypeInt, Optional: true, Description: "type value",
						},
						"type_string": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "domain name",
						},
						"content_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rdata_size_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rdata_all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"record_num_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"record_all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbDnsCacheOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsCacheOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsCacheOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDnsCacheOperOper := setObjectSlbDnsCacheOperOper(res)
		d.Set("oper", SlbDnsCacheOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDnsCacheOperOper(ret edpt.DataSlbDnsCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cache_client":     setSliceSlbDnsCacheOperOperCacheClient(ret.DtSlbDnsCacheOper.Oper.CacheClient),
			"cache_entry":      setSliceSlbDnsCacheOperOperCacheEntry(ret.DtSlbDnsCacheOper.Oper.CacheEntry),
			"total":            ret.DtSlbDnsCacheOper.Oper.Total,
			"client":           ret.DtSlbDnsCacheOper.Oper.Client,
			"entry":            ret.DtSlbDnsCacheOper.Oper.Entry,
			"global":           ret.DtSlbDnsCacheOper.Oper.Global,
			"cache_content":    ret.DtSlbDnsCacheOper.Oper.CacheContent,
			"vport":            ret.DtSlbDnsCacheOper.Oper.Vport,
			"vs_name":          ret.DtSlbDnsCacheOper.Oper.VsName,
			"port_type":        ret.DtSlbDnsCacheOper.Oper.PortType,
			"port_num":         ret.DtSlbDnsCacheOper.Oper.PortNum,
			"type_value":       ret.DtSlbDnsCacheOper.Oper.TypeValue,
			"fqdn_domain":      ret.DtSlbDnsCacheOper.Oper.FqdnDomain,
			"class_string":     ret.DtSlbDnsCacheOper.Oper.ClassString,
			"class_value":      ret.DtSlbDnsCacheOper.Oper.ClassValue,
			"type_string":      ret.DtSlbDnsCacheOper.Oper.TypeString,
			"domain_name":      ret.DtSlbDnsCacheOper.Oper.DomainName,
			"content_mode":     ret.DtSlbDnsCacheOper.Oper.ContentMode,
			"rdata_size_value": ret.DtSlbDnsCacheOper.Oper.RdataSizeValue,
			"rdata_all":        ret.DtSlbDnsCacheOper.Oper.RdataAll,
			"record_num_value": ret.DtSlbDnsCacheOper.Oper.RecordNumValue,
			"record_all":       ret.DtSlbDnsCacheOper.Oper.RecordAll,
		},
	}
}

func setSliceSlbDnsCacheOperOperCacheClient(d []edpt.SlbDnsCacheOperOperCacheClient) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["domain"] = item.Domain
		in["address"] = item.Address
		in["unit_type"] = item.Unit_type
		in["curr_rate"] = item.Curr_rate
		in["over_rate_limit_times"] = item.Over_rate_limit_times
		in["lockup"] = item.Lockup
		in["lockup_time"] = item.Lockup_time
		result = append(result, in)
	}
	return result
}

func setSliceSlbDnsCacheOperOperCacheEntry(d []edpt.SlbDnsCacheOperOperCacheEntry) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["domain"] = item.Domain
		in["dnssec"] = item.Dnssec
		in["cache_negative"] = item.Cache_negative
		in["cache_type"] = item.Cache_type
		in["cache_class"] = item.Cache_class
		in["q_length"] = item.Q_length
		in["r_length"] = item.R_length
		in["ttl"] = item.Ttl
		in["age"] = item.Age
		in["weight"] = item.Weight
		in["hits"] = item.Hits
		in["an_count"] = item.An_count
		in["ns_count"] = item.Ns_count
		in["ar_count"] = item.Ar_count
		in["entry_record"] = setSliceSlbDnsCacheOperOperCacheEntryEntryRecord(item.EntryRecord)
		result = append(result, in)
	}
	return result
}

func setSliceSlbDnsCacheOperOperCacheEntryEntryRecord(d []edpt.SlbDnsCacheOperOperCacheEntryEntryRecord) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["record_type"] = item.Record_type
		in["record_class"] = item.Record_class
		in["record_ttl"] = item.Record_ttl
		in["record_rdlen"] = item.Record_rdlen
		in["record_rdata"] = item.Record_rdata
		in["record_rdata_tc"] = item.Record_rdata_tc
		result = append(result, in)
	}
	return result
}

func getObjectSlbDnsCacheOperOper(d []interface{}) edpt.SlbDnsCacheOperOper {

	count1 := len(d)
	var ret edpt.SlbDnsCacheOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheClient = getSliceSlbDnsCacheOperOperCacheClient(in["cache_client"].([]interface{}))
		ret.CacheEntry = getSliceSlbDnsCacheOperOperCacheEntry(in["cache_entry"].([]interface{}))
		ret.Total = in["total"].(int)
		ret.Client = in["client"].(int)
		ret.Entry = in["entry"].(int)
		ret.Global = in["global"].(int)
		ret.CacheContent = in["cache_content"].(int)
		ret.Vport = in["vport"].(int)
		ret.VsName = in["vs_name"].(string)
		ret.PortType = in["port_type"].(string)
		ret.PortNum = in["port_num"].(int)
		ret.TypeValue = in["type_value"].(int)
		ret.FqdnDomain = in["fqdn_domain"].(string)
		ret.ClassString = in["class_string"].(string)
		ret.ClassValue = in["class_value"].(int)
		ret.TypeString = in["type_string"].(string)
		ret.DomainName = in["domain_name"].(string)
		ret.ContentMode = in["content_mode"].(string)
		ret.RdataSizeValue = in["rdata_size_value"].(int)
		ret.RdataAll = in["rdata_all"].(int)
		ret.RecordNumValue = in["record_num_value"].(int)
		ret.RecordAll = in["record_all"].(int)
	}
	return ret
}

func getSliceSlbDnsCacheOperOperCacheClient(d []interface{}) []edpt.SlbDnsCacheOperOperCacheClient {

	count1 := len(d)
	ret := make([]edpt.SlbDnsCacheOperOperCacheClient, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsCacheOperOperCacheClient
		oi.Domain = in["domain"].(string)
		oi.Address = in["address"].(string)
		oi.Unit_type = in["unit_type"].(string)
		oi.Curr_rate = in["curr_rate"].(int)
		oi.Over_rate_limit_times = in["over_rate_limit_times"].(int)
		oi.Lockup = in["lockup"].(int)
		oi.Lockup_time = in["lockup_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbDnsCacheOperOperCacheEntry(d []interface{}) []edpt.SlbDnsCacheOperOperCacheEntry {

	count1 := len(d)
	ret := make([]edpt.SlbDnsCacheOperOperCacheEntry, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsCacheOperOperCacheEntry
		oi.Name = in["name"].(string)
		oi.Domain = in["domain"].(string)
		oi.Dnssec = in["dnssec"].(int)
		oi.Cache_negative = in["cache_negative"].(int)
		oi.Cache_type = in["cache_type"].(int)
		oi.Cache_class = in["cache_class"].(int)
		oi.Q_length = in["q_length"].(int)
		oi.R_length = in["r_length"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.Age = in["age"].(int)
		oi.Weight = in["weight"].(int)
		oi.Hits = in["hits"].(int)
		oi.An_count = in["an_count"].(int)
		oi.Ns_count = in["ns_count"].(int)
		oi.Ar_count = in["ar_count"].(int)
		oi.EntryRecord = getSliceSlbDnsCacheOperOperCacheEntryEntryRecord(in["entry_record"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbDnsCacheOperOperCacheEntryEntryRecord(d []interface{}) []edpt.SlbDnsCacheOperOperCacheEntryEntryRecord {

	count1 := len(d)
	ret := make([]edpt.SlbDnsCacheOperOperCacheEntryEntryRecord, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsCacheOperOperCacheEntryEntryRecord
		oi.Record_type = in["record_type"].(int)
		oi.Record_class = in["record_class"].(int)
		oi.Record_ttl = in["record_ttl"].(int)
		oi.Record_rdlen = in["record_rdlen"].(int)
		oi.Record_rdata = in["record_rdata"].(string)
		oi.Record_rdata_tc = in["record_rdata_tc"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDnsCacheOper(d *schema.ResourceData) edpt.SlbDnsCacheOper {
	var ret edpt.SlbDnsCacheOper

	ret.Oper = getObjectSlbDnsCacheOperOper(d.Get("oper").([]interface{}))
	return ret
}
