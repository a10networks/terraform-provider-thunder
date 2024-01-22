package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_zone_stats`: Statistics for the object zone\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbZoneStatsRead,

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
										Type: schema.TypeInt, Optional: true, Description: "Number of times the record has been used",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify the name for the DNS zone",
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
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"received_query": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of DNS queries received for the zone",
						},
						"sent_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of DNS replies sent to clients for the zone",
						},
						"proxy_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of DNS replies sent to clients by the ACOS device as a DNS proxy for the zone",
						},
						"cache_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of cached DNS replies sent to clients by the ACOS device for the zone. (This statistic applies only if the DNS cac",
						},
						"server_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of DNS replies sent to clients by the ACOS device as a DNS server for the zone. (This statistic applies only if th",
						},
						"sticky_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of DNS replies sent to clients by the ACOS device to keep the clients on the same site. (This statistic applies on",
						},
						"backup_mode_response": {
							Type: schema.TypeInt, Optional: true, Description: "Total Number of DNS replies sent to clients by the ACOS device in backup mode",
						},
					},
				},
			},
		},
	}
}

func resourceGslbZoneStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbZoneStatsDnsCaaRecordList := setSliceGslbZoneStatsDnsCaaRecordList(res)
		d.Set("dns_caa_record_list", GslbZoneStatsDnsCaaRecordList)
		GslbZoneStatsDnsMxRecordList := setSliceGslbZoneStatsDnsMxRecordList(res)
		d.Set("dns_mx_record_list", GslbZoneStatsDnsMxRecordList)
		GslbZoneStatsDnsNsRecordList := setSliceGslbZoneStatsDnsNsRecordList(res)
		d.Set("dns_ns_record_list", GslbZoneStatsDnsNsRecordList)
		GslbZoneStatsServiceList := setSliceGslbZoneStatsServiceList(res)
		d.Set("service_list", GslbZoneStatsServiceList)
		GslbZoneStatsStats := setObjectGslbZoneStatsStats(res)
		d.Set("stats", GslbZoneStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbZoneStatsDnsCaaRecordList(d edpt.DataGslbZoneStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneStats.DnsCaaRecordList {
		in := make(map[string]interface{})
		in["critical_flag"] = item.CriticalFlag
		in["property_tag"] = item.PropertyTag
		in["rdata"] = item.Rdata
		in["stats"] = setObjectGslbZoneStatsDnsCaaRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsDnsCaaRecordListStats(d edpt.GslbZoneStatsDnsCaaRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsDnsMxRecordList(d edpt.DataGslbZoneStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneStats.DnsMxRecordList {
		in := make(map[string]interface{})
		in["mx_name"] = item.MxName
		in["stats"] = setObjectGslbZoneStatsDnsMxRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsDnsMxRecordListStats(d edpt.GslbZoneStatsDnsMxRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsDnsNsRecordList(d edpt.DataGslbZoneStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneStats.DnsNsRecordList {
		in := make(map[string]interface{})
		in["ns_name"] = item.NsName
		in["stats"] = setObjectGslbZoneStatsDnsNsRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsDnsNsRecordListStats(d edpt.GslbZoneStatsDnsNsRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceList(d edpt.DataGslbZoneStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbZoneStats.ServiceList {
		in := make(map[string]interface{})
		in["service_port"] = item.ServicePort
		in["service_name"] = item.ServiceName
		in["stats"] = setObjectGslbZoneStatsServiceListStats(item.Stats)
		in["dns_cname_record_list"] = setSliceGslbZoneStatsServiceListDnsCnameRecordList(item.DnsCnameRecordList)
		in["dns_mx_record_list"] = setSliceGslbZoneStatsServiceListDnsMxRecordList(item.DnsMxRecordList)
		in["dns_ns_record_list"] = setSliceGslbZoneStatsServiceListDnsNsRecordList(item.DnsNsRecordList)
		in["dns_ptr_record_list"] = setSliceGslbZoneStatsServiceListDnsPtrRecordList(item.DnsPtrRecordList)
		in["dns_srv_record_list"] = setSliceGslbZoneStatsServiceListDnsSrvRecordList(item.DnsSrvRecordList)
		in["dns_naptr_record_list"] = setSliceGslbZoneStatsServiceListDnsNaptrRecordList(item.DnsNaptrRecordList)
		in["dns_txt_record_list"] = setSliceGslbZoneStatsServiceListDnsTxtRecordList(item.DnsTxtRecordList)
		in["dns_caa_record_list"] = setSliceGslbZoneStatsServiceListDnsCaaRecordList(item.DnsCaaRecordList)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListStats(d edpt.GslbZoneStatsServiceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["received_query"] = d.ReceivedQuery

	in["sent_response"] = d.SentResponse

	in["proxy_mode_response"] = d.ProxyModeResponse

	in["cache_mode_response"] = d.CacheModeResponse

	in["server_mode_response"] = d.ServerModeResponse

	in["sticky_mode_response"] = d.StickyModeResponse

	in["backup_mode_response"] = d.BackupModeResponse
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsCnameRecordList(d []edpt.GslbZoneStatsServiceListDnsCnameRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["alias_name"] = item.AliasName
		in["stats"] = setObjectGslbZoneStatsServiceListDnsCnameRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsCnameRecordListStats(d edpt.GslbZoneStatsServiceListDnsCnameRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["cname_hits"] = d.CnameHits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsMxRecordList(d []edpt.GslbZoneStatsServiceListDnsMxRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["mx_name"] = item.MxName
		in["stats"] = setObjectGslbZoneStatsServiceListDnsMxRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsMxRecordListStats(d edpt.GslbZoneStatsServiceListDnsMxRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsNsRecordList(d []edpt.GslbZoneStatsServiceListDnsNsRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ns_name"] = item.NsName
		in["stats"] = setObjectGslbZoneStatsServiceListDnsNsRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsNsRecordListStats(d edpt.GslbZoneStatsServiceListDnsNsRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsPtrRecordList(d []edpt.GslbZoneStatsServiceListDnsPtrRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ptr_name"] = item.PtrName
		in["stats"] = setObjectGslbZoneStatsServiceListDnsPtrRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsPtrRecordListStats(d edpt.GslbZoneStatsServiceListDnsPtrRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsSrvRecordList(d []edpt.GslbZoneStatsServiceListDnsSrvRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["srv_name"] = item.SrvName
		in["port"] = item.Port
		in["stats"] = setObjectGslbZoneStatsServiceListDnsSrvRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsSrvRecordListStats(d edpt.GslbZoneStatsServiceListDnsSrvRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsNaptrRecordList(d []edpt.GslbZoneStatsServiceListDnsNaptrRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["naptr_target"] = item.NaptrTarget
		in["service_proto"] = item.ServiceProto
		in["flag"] = item.Flag
		in["stats"] = setObjectGslbZoneStatsServiceListDnsNaptrRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsNaptrRecordListStats(d edpt.GslbZoneStatsServiceListDnsNaptrRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["naptr_hits"] = d.NaptrHits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsTxtRecordList(d []edpt.GslbZoneStatsServiceListDnsTxtRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["record_name"] = item.RecordName
		in["stats"] = setObjectGslbZoneStatsServiceListDnsTxtRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsTxtRecordListStats(d edpt.GslbZoneStatsServiceListDnsTxtRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setSliceGslbZoneStatsServiceListDnsCaaRecordList(d []edpt.GslbZoneStatsServiceListDnsCaaRecordList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["critical_flag"] = item.CriticalFlag
		in["property_tag"] = item.PropertyTag
		in["rdata"] = item.Rdata
		in["stats"] = setObjectGslbZoneStatsServiceListDnsCaaRecordListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbZoneStatsServiceListDnsCaaRecordListStats(d edpt.GslbZoneStatsServiceListDnsCaaRecordListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hits"] = d.Hits
	result = append(result, in)
	return result
}

func setObjectGslbZoneStatsStats(ret edpt.DataGslbZoneStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"received_query":       ret.DtGslbZoneStats.Stats.ReceivedQuery,
			"sent_response":        ret.DtGslbZoneStats.Stats.SentResponse,
			"proxy_mode_response":  ret.DtGslbZoneStats.Stats.ProxyModeResponse,
			"cache_mode_response":  ret.DtGslbZoneStats.Stats.CacheModeResponse,
			"server_mode_response": ret.DtGslbZoneStats.Stats.ServerModeResponse,
			"sticky_mode_response": ret.DtGslbZoneStats.Stats.StickyModeResponse,
			"backup_mode_response": ret.DtGslbZoneStats.Stats.BackupModeResponse,
		},
	}
}

func getSliceGslbZoneStatsDnsCaaRecordList(d []interface{}) []edpt.GslbZoneStatsDnsCaaRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsDnsCaaRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsDnsCaaRecordList
		oi.CriticalFlag = in["critical_flag"].(int)
		oi.PropertyTag = in["property_tag"].(string)
		oi.Rdata = in["rdata"].(string)
		oi.Stats = getObjectGslbZoneStatsDnsCaaRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsDnsCaaRecordListStats(d []interface{}) edpt.GslbZoneStatsDnsCaaRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsDnsCaaRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsDnsMxRecordList(d []interface{}) []edpt.GslbZoneStatsDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Stats = getObjectGslbZoneStatsDnsMxRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsDnsMxRecordListStats(d []interface{}) edpt.GslbZoneStatsDnsMxRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsDnsMxRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsDnsNsRecordList(d []interface{}) []edpt.GslbZoneStatsDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Stats = getObjectGslbZoneStatsDnsNsRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsDnsNsRecordListStats(d []interface{}) edpt.GslbZoneStatsDnsNsRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsDnsNsRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceList(d []interface{}) []edpt.GslbZoneStatsServiceList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceList
		oi.ServicePort = in["service_port"].(int)
		oi.ServiceName = in["service_name"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListStats(in["stats"].([]interface{}))
		oi.DnsCnameRecordList = getSliceGslbZoneStatsServiceListDnsCnameRecordList(in["dns_cname_record_list"].([]interface{}))
		oi.DnsMxRecordList = getSliceGslbZoneStatsServiceListDnsMxRecordList(in["dns_mx_record_list"].([]interface{}))
		oi.DnsNsRecordList = getSliceGslbZoneStatsServiceListDnsNsRecordList(in["dns_ns_record_list"].([]interface{}))
		oi.DnsPtrRecordList = getSliceGslbZoneStatsServiceListDnsPtrRecordList(in["dns_ptr_record_list"].([]interface{}))
		oi.DnsSrvRecordList = getSliceGslbZoneStatsServiceListDnsSrvRecordList(in["dns_srv_record_list"].([]interface{}))
		oi.DnsNaptrRecordList = getSliceGslbZoneStatsServiceListDnsNaptrRecordList(in["dns_naptr_record_list"].([]interface{}))
		oi.DnsTxtRecordList = getSliceGslbZoneStatsServiceListDnsTxtRecordList(in["dns_txt_record_list"].([]interface{}))
		oi.DnsCaaRecordList = getSliceGslbZoneStatsServiceListDnsCaaRecordList(in["dns_caa_record_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListStats(d []interface{}) edpt.GslbZoneStatsServiceListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListStats
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

func getSliceGslbZoneStatsServiceListDnsCnameRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsCnameRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsCnameRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsCnameRecordList
		oi.AliasName = in["alias_name"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsCnameRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsCnameRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsCnameRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsCnameRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CnameHits = in["cname_hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsMxRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsMxRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsMxRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsMxRecordList
		oi.MxName = in["mx_name"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsMxRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsMxRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsMxRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsMxRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsNsRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsNsRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsNsRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsNsRecordList
		oi.NsName = in["ns_name"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsNsRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsNsRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsNsRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsNsRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsPtrRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsPtrRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsPtrRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsPtrRecordList
		oi.PtrName = in["ptr_name"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsPtrRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsPtrRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsPtrRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsPtrRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsSrvRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsSrvRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsSrvRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsSrvRecordList
		oi.SrvName = in["srv_name"].(string)
		oi.Port = in["port"].(int)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsSrvRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsSrvRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsSrvRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsSrvRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsNaptrRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsNaptrRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsNaptrRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsNaptrRecordList
		oi.NaptrTarget = in["naptr_target"].(string)
		oi.ServiceProto = in["service_proto"].(string)
		oi.Flag = in["flag"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsNaptrRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsNaptrRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsNaptrRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsNaptrRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NaptrHits = in["naptr_hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsTxtRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsTxtRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsTxtRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsTxtRecordList
		oi.RecordName = in["record_name"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsTxtRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsTxtRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsTxtRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsTxtRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getSliceGslbZoneStatsServiceListDnsCaaRecordList(d []interface{}) []edpt.GslbZoneStatsServiceListDnsCaaRecordList {

	count1 := len(d)
	ret := make([]edpt.GslbZoneStatsServiceListDnsCaaRecordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneStatsServiceListDnsCaaRecordList
		oi.CriticalFlag = in["critical_flag"].(int)
		oi.PropertyTag = in["property_tag"].(string)
		oi.Rdata = in["rdata"].(string)
		oi.Stats = getObjectGslbZoneStatsServiceListDnsCaaRecordListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbZoneStatsServiceListDnsCaaRecordListStats(d []interface{}) edpt.GslbZoneStatsServiceListDnsCaaRecordListStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsServiceListDnsCaaRecordListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func getObjectGslbZoneStatsStats(d []interface{}) edpt.GslbZoneStatsStats {

	count1 := len(d)
	var ret edpt.GslbZoneStatsStats
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

func dataToEndpointGslbZoneStats(d *schema.ResourceData) edpt.GslbZoneStats {
	var ret edpt.GslbZoneStats

	ret.DnsCaaRecordList = getSliceGslbZoneStatsDnsCaaRecordList(d.Get("dns_caa_record_list").([]interface{}))

	ret.DnsMxRecordList = getSliceGslbZoneStatsDnsMxRecordList(d.Get("dns_mx_record_list").([]interface{}))

	ret.DnsNsRecordList = getSliceGslbZoneStatsDnsNsRecordList(d.Get("dns_ns_record_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.ServiceList = getSliceGslbZoneStatsServiceList(d.Get("service_list").([]interface{}))

	ret.Stats = getObjectGslbZoneStatsStats(d.Get("stats").([]interface{}))
	return ret
}
