package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDnsCacheOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_dns_cache_oper`: Operational Status for the object dns-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDnsCacheOperRead,

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

func resourceSystemDnsCacheOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCacheOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsCacheOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDnsCacheOperOper := setObjectSystemDnsCacheOperOper(res)
		d.Set("oper", SystemDnsCacheOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemDnsCacheOperOper(ret edpt.DataSystemDnsCacheOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cache_client":     setSliceSystemDnsCacheOperOperCacheClient(ret.DtSystemDnsCacheOper.Oper.CacheClient),
			"cache_entry":      setSliceSystemDnsCacheOperOperCacheEntry(ret.DtSystemDnsCacheOper.Oper.CacheEntry),
			"total":            ret.DtSystemDnsCacheOper.Oper.Total,
			"client":           ret.DtSystemDnsCacheOper.Oper.Client,
			"entry":            ret.DtSystemDnsCacheOper.Oper.Entry,
			"global":           ret.DtSystemDnsCacheOper.Oper.Global,
			"cache_content":    ret.DtSystemDnsCacheOper.Oper.CacheContent,
			"vport":            ret.DtSystemDnsCacheOper.Oper.Vport,
			"vs_name":          ret.DtSystemDnsCacheOper.Oper.VsName,
			"port_type":        ret.DtSystemDnsCacheOper.Oper.PortType,
			"port_num":         ret.DtSystemDnsCacheOper.Oper.PortNum,
			"type_value":       ret.DtSystemDnsCacheOper.Oper.TypeValue,
			"fqdn_domain":      ret.DtSystemDnsCacheOper.Oper.FqdnDomain,
			"class_string":     ret.DtSystemDnsCacheOper.Oper.ClassString,
			"class_value":      ret.DtSystemDnsCacheOper.Oper.ClassValue,
			"type_string":      ret.DtSystemDnsCacheOper.Oper.TypeString,
			"domain_name":      ret.DtSystemDnsCacheOper.Oper.DomainName,
			"content_mode":     ret.DtSystemDnsCacheOper.Oper.ContentMode,
			"rdata_size_value": ret.DtSystemDnsCacheOper.Oper.RdataSizeValue,
			"rdata_all":        ret.DtSystemDnsCacheOper.Oper.RdataAll,
			"record_num_value": ret.DtSystemDnsCacheOper.Oper.RecordNumValue,
			"record_all":       ret.DtSystemDnsCacheOper.Oper.RecordAll,
		},
	}
}

func setSliceSystemDnsCacheOperOperCacheClient(d []edpt.SystemDnsCacheOperOperCacheClient) []map[string]interface{} {
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

func setSliceSystemDnsCacheOperOperCacheEntry(d []edpt.SystemDnsCacheOperOperCacheEntry) []map[string]interface{} {
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
		in["entry_record"] = setSliceSystemDnsCacheOperOperCacheEntryEntryRecord(item.EntryRecord)
		result = append(result, in)
	}
	return result
}

func setSliceSystemDnsCacheOperOperCacheEntryEntryRecord(d []edpt.SystemDnsCacheOperOperCacheEntryEntryRecord) []map[string]interface{} {
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

func getObjectSystemDnsCacheOperOper(d []interface{}) edpt.SystemDnsCacheOperOper {

	count1 := len(d)
	var ret edpt.SystemDnsCacheOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CacheClient = getSliceSystemDnsCacheOperOperCacheClient(in["cache_client"].([]interface{}))
		ret.CacheEntry = getSliceSystemDnsCacheOperOperCacheEntry(in["cache_entry"].([]interface{}))
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

func getSliceSystemDnsCacheOperOperCacheClient(d []interface{}) []edpt.SystemDnsCacheOperOperCacheClient {

	count1 := len(d)
	ret := make([]edpt.SystemDnsCacheOperOperCacheClient, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsCacheOperOperCacheClient
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

func getSliceSystemDnsCacheOperOperCacheEntry(d []interface{}) []edpt.SystemDnsCacheOperOperCacheEntry {

	count1 := len(d)
	ret := make([]edpt.SystemDnsCacheOperOperCacheEntry, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsCacheOperOperCacheEntry
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
		oi.EntryRecord = getSliceSystemDnsCacheOperOperCacheEntryEntryRecord(in["entry_record"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemDnsCacheOperOperCacheEntryEntryRecord(d []interface{}) []edpt.SystemDnsCacheOperOperCacheEntryEntryRecord {

	count1 := len(d)
	ret := make([]edpt.SystemDnsCacheOperOperCacheEntryEntryRecord, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsCacheOperOperCacheEntryEntryRecord
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

func dataToEndpointSystemDnsCacheOper(d *schema.ResourceData) edpt.SystemDnsCacheOper {
	var ret edpt.SystemDnsCacheOper

	ret.Oper = getObjectSystemDnsCacheOperOper(d.Get("oper").([]interface{}))
	return ret
}
