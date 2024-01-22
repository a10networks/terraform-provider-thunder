package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_service_stats`: Statistics for the object service\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneServiceStatsRead,

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
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the CAA has been used",
									},
								},
							},
						},
					},
				},
			},
			"dns_cname_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the alias name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cname_hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the CNAME has been used",
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
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
									},
								},
							},
						},
					},
				},
			},
			"dns_naptr_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"naptr_target": {
							Type: schema.TypeString, Required: true, Description: "Specify the replacement or regular expression",
						},
						"service_proto": {
							Type: schema.TypeString, Required: true, Description: "Specify Service and Protocol",
						},
						"flag": {
							Type: schema.TypeString, Required: true, Description: "Specify the flag (e.g., a, s). Default is empty flag",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"naptr_hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the NAPTR has been used",
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
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
									},
								},
							},
						},
					},
				},
			},
			"dns_ptr_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ptr_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
									},
								},
							},
						},
					},
				},
			},
			"dns_srv_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"srv_name": {
							Type: schema.TypeString, Required: true, Description: "Specify Domain Name",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Specify Port (Port Number)",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
									},
								},
							},
						},
					},
				},
			},
			"dns_txt_record_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"record_name": {
							Type: schema.TypeString, Required: true, Description: "Specify the Object Name for TXT Data",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
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
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"received_query": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DNS queries received for the service",
						},
						"sent_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DNS replies sent to clients for the service",
						},
						"proxy_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DNS replies sent to clients by the ACOS device as a DNS proxy for the service",
						},
						"cache_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of cached DNS replies sent to clients by the ACOS device for the service. (This statistic applies only if the DNS cache",
						},
						"server_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DNS replies sent to clients by the ACOS device as a DNS server for the service. (This statistic applies only if the D",
						},
						"sticky_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Number of DNS replies sent to clients by the ACOS device to keep the clients on the same site. (This statistic applies only if",
						},
						"backup_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "help Number of DNS replies sent to clients by the ACOS device in backup mode",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceGslbZoneServiceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneServiceStatsDnsCaaRecordList := setSliceGslbZoneServiceStatsDnsCaaRecordList(res)
		d.Set("dns_caa_record_list", GslbZoneServiceStatsDnsCaaRecordList)
		GslbZoneServiceStatsDnsCnameRecordList := setSliceGslbZoneServiceStatsDnsCnameRecordList(res)
		d.Set("dns_cname_record_list", GslbZoneServiceStatsDnsCnameRecordList)
		GslbZoneServiceStatsDnsMxRecordList := setSliceGslbZoneServiceStatsDnsMxRecordList(res)
		d.Set("dns_mx_record_list", GslbZoneServiceStatsDnsMxRecordList)
		GslbZoneServiceStatsDnsNaptrRecordList := setSliceGslbZoneServiceStatsDnsNaptrRecordList(res)
		d.Set("dns_naptr_record_list", GslbZoneServiceStatsDnsNaptrRecordList)
		GslbZoneServiceStatsDnsNsRecordList := setSliceGslbZoneServiceStatsDnsNsRecordList(res)
		d.Set("dns_ns_record_list", GslbZoneServiceStatsDnsNsRecordList)
		GslbZoneServiceStatsDnsPtrRecordList := setSliceGslbZoneServiceStatsDnsPtrRecordList(res)
		d.Set("dns_ptr_record_list", GslbZoneServiceStatsDnsPtrRecordList)
		GslbZoneServiceStatsDnsSrvRecordList := setSliceGslbZoneServiceStatsDnsSrvRecordList(res)
		d.Set("dns_srv_record_list", GslbZoneServiceStatsDnsSrvRecordList)
		GslbZoneServiceStatsDnsTxtRecordList := setSliceGslbZoneServiceStatsDnsTxtRecordList(res)
		d.Set("dns_txt_record_list", GslbZoneServiceStatsDnsTxtRecordList)
		GslbZoneServiceStatsStats := setObjectGslbZoneServiceStatsStats(res)
		d.Set("stats", GslbZoneServiceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbZoneServiceStatsDnsCaaRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsCaaRecordList {
		in := make(map[string]interface{})
		in["critical_flag"] = item.CriticalFlag
		in["property_tag"] = item.PropertyTag
		in["rdata"] = item.Rdata
		in["stats"] = setObjectGslbZoneServiceStatsDnsCaaRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsCaaRecordListStats(d edpt.GslbZoneServiceStatsDnsCaaRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsCnameRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsCnameRecordList {
		in := make(map[string]interface{})
		in["alias_name"] = item.AliasName
		in["stats"] = setObjectGslbZoneServiceStatsDnsCnameRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsCnameRecordListStats(d edpt.GslbZoneServiceStatsDnsCnameRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["cname_hits"] = d.CnameHits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsMxRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsMxRecordList {
		in := make(map[string]interface{})
		in["mx_name"] = item.MxName
		in["stats"] = setObjectGslbZoneServiceStatsDnsMxRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsMxRecordListStats(d edpt.GslbZoneServiceStatsDnsMxRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsNaptrRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsNaptrRecordList {
		in := make(map[string]interface{})
		in["naptr_target"] = item.NaptrTarget
		in["service_proto"] = item.ServiceProto
		in["flag"] = item.Flag
		in["stats"] = setObjectGslbZoneServiceStatsDnsNaptrRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsNaptrRecordListStats(d edpt.GslbZoneServiceStatsDnsNaptrRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["naptr_hits"] = d.NaptrHits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsNsRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsNsRecordList {
		in := make(map[string]interface{})
		in["ns_name"] = item.NsName
		in["stats"] = setObjectGslbZoneServiceStatsDnsNsRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsNsRecordListStats(d edpt.GslbZoneServiceStatsDnsNsRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsPtrRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsPtrRecordList {
		in := make(map[string]interface{})
		in["ptr_name"] = item.PtrName
		in["stats"] = setObjectGslbZoneServiceStatsDnsPtrRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsPtrRecordListStats(d edpt.GslbZoneServiceStatsDnsPtrRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsSrvRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsSrvRecordList {
		in := make(map[string]interface{})
		in["srv_name"] = item.SrvName
		in["port"] = item.Port
		in["stats"] = setObjectGslbZoneServiceStatsDnsSrvRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsSrvRecordListStats(d edpt.GslbZoneServiceStatsDnsSrvRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneServiceStatsDnsTxtRecordList(d edpt.DataGslbZoneServiceStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneServiceStats.DnsTxtRecordList {
		in := make(map[string]interface{})
		in["record_name"] = item.RecordName
		in["stats"] = setObjectGslbZoneServiceStatsDnsTxtRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneServiceStatsDnsTxtRecordListStats(d edpt.GslbZoneServiceStatsDnsTxtRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setObjectGslbZoneServiceStatsStats(ret edpt.DataGslbZoneServiceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"received_query":       ret.DtGslbZoneServiceStats.Stats.ReceivedQuery,
			"sent_response":        ret.DtGslbZoneServiceStats.Stats.SentResponse,
			"proxy_mode_response":  ret.DtGslbZoneServiceStats.Stats.ProxyModeResponse,
			"cache_mode_response":  ret.DtGslbZoneServiceStats.Stats.CacheModeResponse,
			"server_mode_response": ret.DtGslbZoneServiceStats.Stats.ServerModeResponse,
			"sticky_mode_response": ret.DtGslbZoneServiceStats.Stats.StickyModeResponse,
			"backup_mode_response": ret.DtGslbZoneServiceStats.Stats.BackupModeResponse,
		},
	}
}

func getSliceGslbZoneServiceStatsDnsCaaRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsCaaRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsCaaRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsCaaRecordList
		oi.CriticalFlag = in["critical_flag"].(int)
		oi.PropertyTag = in["property_tag"].(string)
		oi.Rdata = in["rdata"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsCaaRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsCaaRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsCaaRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsCaaRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsCnameRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsCnameRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsCnameRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsCnameRecordList
		oi.AliasName = in["alias_name"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsCnameRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsCnameRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsCnameRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsCnameRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CnameHits = in["cname_hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsMxRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsMxRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsMxRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsMxRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsMxRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsNaptrRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsNaptrRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsNaptrRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsNaptrRecordList
		oi.NaptrTarget = in["naptr_target"].(string)
		oi.ServiceProto = in["service_proto"].(string)
		oi.Flag = in["flag"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsNaptrRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsNaptrRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsNaptrRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsNaptrRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NaptrHits = in["naptr_hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsNsRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsNsRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsNsRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsNsRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsNsRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsPtrRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsPtrRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsPtrRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsPtrRecordList
		oi.PtrName = in["ptr_name"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsPtrRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsPtrRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsPtrRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsPtrRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsSrvRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsSrvRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsSrvRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsSrvRecordList
		oi.SrvName = in["srv_name"].(string)
		oi.Port = in["port"].(int)
		oi.Stats = getObjectGslbZoneServiceStatsDnsSrvRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsSrvRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsSrvRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsSrvRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneServiceStatsDnsTxtRecordList(d []interface{}) []edpt.GslbZoneServiceStatsDnsTxtRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceStatsDnsTxtRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceStatsDnsTxtRecordList
		oi.RecordName = in["record_name"].(string)
		oi.Stats = getObjectGslbZoneServiceStatsDnsTxtRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneServiceStatsDnsTxtRecordListStats(d []interface{}) edpt.GslbZoneServiceStatsDnsTxtRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsDnsTxtRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getObjectGslbZoneServiceStatsStats(d []interface{}) edpt.GslbZoneServiceStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneServiceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ReceivedQuery = in["received_query"].(int)
		ret.SentResponse = in["sent_response"].(int)
		ret.ProxyModeResponse = in["proxy_mode_response"].(int)
		ret.CacheModeResponse = in["cache_mode_response"].(int)
		ret.ServerModeResponse = in["server_mode_response"].(int)
		ret.StickyModeResponse = in["sticky_mode_response"].(int)
		ret.BackupModeResponse = in["backup_mode_response"].(int)
	}
	return ret
}

func dataToEndpointGslbZoneServiceStats(d *schema.ResourceData) edpt.GslbZoneServiceStats {
	var ret edpt.GslbZoneServiceStats

	ret.DnsCaaRecordList = getSliceGslbZoneServiceStatsDnsCaaRecordList(d.Get("dns_caa_record_list").([]interface{}))

	ret.DnsCnameRecordList = getSliceGslbZoneServiceStatsDnsCnameRecordList(d.Get("dns_cname_record_list").([]interface{}))

	ret.DnsMxRecordList = getSliceGslbZoneServiceStatsDnsMxRecordList(d.Get("dns_mx_record_list").([]interface{}))

	ret.DnsNaptrRecordList = getSliceGslbZoneServiceStatsDnsNaptrRecordList(d.Get("dns_naptr_record_list").([]interface{}))

	ret.DnsNsRecordList = getSliceGslbZoneServiceStatsDnsNsRecordList(d.Get("dns_ns_record_list").([]interface{}))

	ret.DnsPtrRecordList = getSliceGslbZoneServiceStatsDnsPtrRecordList(d.Get("dns_ptr_record_list").([]interface{}))

	ret.DnsSrvRecordList = getSliceGslbZoneServiceStatsDnsSrvRecordList(d.Get("dns_srv_record_list").([]interface{}))

	ret.DnsTxtRecordList = getSliceGslbZoneServiceStatsDnsTxtRecordList(d.Get("dns_txt_record_list").([]interface{}))

	ret.ServiceName = d.Get("service_name").(string)

	ret.ServicePort = d.Get("service_port").(int)

	ret.Stats = getObjectGslbZoneServiceStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
